package data

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"sync"
	"syscall"

	lru "github.com/hashicorp/golang-lru"
	"github.com/tsandl/skvdb/leveldb/storage"
)

var errFileOpen = errors.New("leveldb/storage: file still open")

const CacheSize = 500

type OpenOption struct {
	Bucket        string
	Path          string
	Ak            string
	Sk            string
	Region        string
	Endpoint      string
	LocalCacheDir string
}

type S3StorageLock struct {
	ms *S3Storage
}

func (lock *S3StorageLock) Unlock() {
	ms := lock.ms
	ms.objStore.Remove("LOCK")
	if ms.slock == lock {
		ms.slock = nil
	}
	return
}

// S3Storage is a s3-backed storage.
type S3Storage struct {
	mu       sync.Mutex
	slock    *S3StorageLock
	meta     storage.FileDesc
	objStore *S3Client
	ramFiles *lru.Cache
	opt      OpenOption
}

// NewS3Storage 返回一个基于S3接口的存储实现
func NewS3Storage(opt OpenOption) (storage.Storage, error) {
	s3Client, err := GetS3Client(opt)
	if err != nil {
		return nil, err
	}
	if opt.LocalCacheDir != "" {
		opt.LocalCacheDir = path.Join(opt.LocalCacheDir, opt.Path)
	} else {
		return nil, errors.New("need a local cache dir for s3 storage")
	}
	ramFileCache, _ := lru.New(CacheSize)
	ms := &S3Storage{
		objStore: s3Client,
		ramFiles: ramFileCache,
		opt:      opt,
	}
	return ms, nil
}

// 上传文件
func (ms *S3Storage) uploadFiles() error {
	logFiles, _ := os.ReadDir(ms.opt.LocalCacheDir)
	for _, logF := range logFiles {
		if logF.Name() == "LOCK" {
			continue
		}
		fullName := path.Join(ms.opt.LocalCacheDir, logF.Name())
		content, err := os.ReadFile(fullName)
		if err != nil {
			return err
		}
		fd, _ := fsParseName(logF.Name())
		err = ms.objStore.PutBytes(fd.String(), content)
		if err != nil {
			return err
		}
		os.Remove(fullName)
	}
	return nil
}

// 锁
func (ms *S3Storage) Lock() (storage.Locker, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if ms.slock != nil {
		return nil, storage.ErrLocked
	}
	cacheDir, _ := os.ReadDir(ms.opt.LocalCacheDir)
	if len(cacheDir) > 0 {
		current, _ := ms.objStore.GetBytes("CURRENT")
		if len(current) == 0 {
			log.Println("remove dirty cache files")
			os.RemoveAll(ms.opt.LocalCacheDir)
		}
	}
	err := os.MkdirAll(ms.opt.LocalCacheDir, 0755)
	if err != nil {
		log.Println("LOCK fail", err)
		return nil, storage.ErrLocked
	}
	locked, _ := ms.objStore.GetBytes("LOCK")
	lockFileName := path.Join(ms.opt.LocalCacheDir, "LOCK")
	localSession, _ := os.ReadFile(lockFileName)
	if string(locked) != string(localSession) && len(locked) > 0 {
		log.Println("miss match", string(locked), string(localSession))
		return nil, storage.ErrLocked
	}
	if len(localSession) > 0 {
		var owerPid, rnd int
		fmt.Sscanf(string(localSession), "%d %d", &rnd, &owerPid)
		selfPid := os.Getpid()
		if owerPid != selfPid && syscall.Kill(owerPid, 0) == nil {
			log.Println("locked by another process", "pid", owerPid)
			return nil, storage.ErrLocked
		}
	}
	newSession := fmt.Sprintf("%d %d", rand.Int(), os.Getpid())
	err = ms.objStore.PutBytes("LOCK", []byte(newSession))
	if err != nil {
		log.Println("LOCK fail", err)
		return nil, storage.ErrLocked
	}
	err = os.WriteFile(lockFileName, []byte(newSession), 0644)
	if err != nil {
		log.Println("LOCK fail", err)
		return nil, storage.ErrLocked
	}
	ms.slock = &S3StorageLock{ms: ms}
	err = ms.uploadFiles()
	if err != nil {
		return nil, err
	}
	return ms.slock, nil
}

func (*S3Storage) Log(str string) {}

// 设置元数据
func (ms *S3Storage) SetMeta(fd storage.FileDesc) error {
	if !storage.FileDescOk(fd) {
		return storage.ErrInvalidFile
	}

	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.meta = fd
	metaStr := fmt.Sprintf("%d %d", fd.Type, int64(fd.Num))
	err := ms.objStore.PutBytes("CURRENT", []byte(metaStr))
	if err != nil {
		return err
	}
	return nil
}

// 获取元数据
func (ms *S3Storage) GetMeta() (storage.FileDesc, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if ms.meta.Zero() {
		metaStr, err := ms.objStore.GetBytes("CURRENT")
		if err != nil {
			return storage.FileDesc{}, os.ErrNotExist
		}
		var fType int
		var fNum int64
		fmt.Sscanf(string(metaStr), "%d %d", &fType, &fNum)
		meta := storage.FileDesc{Type: storage.FileType(fType), Num: fNum}
		log.Println("GetMeta Remote", meta)
		ms.meta = meta
	}
	if ms.meta.Zero() {
		return storage.FileDesc{}, os.ErrNotExist
	}
	return ms.meta, nil
}

// 列举 storage中文件信息
func (ms *S3Storage) List(ft storage.FileType) ([]storage.FileDesc, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	fdsAll, err := ms.objStore.List()
	if err != nil {
		return nil, err
	}
	var fds []storage.FileDesc
	for _, fd := range fdsAll {
		if fd.Type&ft != 0 {
			fds = append(fds, fd)
		}
	}
	return fds, nil
}

// 根据FileDesc打开blockserver与minio的联系，读数据
func (ms *S3Storage) Open(fd storage.FileDesc) (storage.Reader, error) {
	if !storage.FileDescOk(fd) {
		return nil, storage.ErrInvalidFile
	}

	ms.mu.Lock()
	defer ms.mu.Unlock()
	if mfileObj, ok := ms.ramFiles.Get(fd.String()); ok {
		mfile := mfileObj.(*memFile)
		mfile.open = true
		return &memReader{Reader: bytes.NewReader(mfile.Bytes()), ms: ms, m: mfile, fd: fd}, nil
	}
	fname := fd.String()
	data, err := ms.objStore.GetBytes(fname)
	if err != nil {
		data, err = os.ReadFile(path.Join(ms.opt.LocalCacheDir, fd.String()))
		if err != nil {
			return nil, os.ErrNotExist
		}
	}
	m := &memFile{Buffer: *bytes.NewBuffer(data), open: true}
	ms.ramFiles.Add(fd.String(), m)
	return &memReader{Reader: bytes.NewReader(data), ms: ms, m: m, fd: fd}, nil
}

// 向minio中写入数据
func (ms *S3Storage) Create(fd storage.FileDesc) (storage.Writer, error) {
	if !storage.FileDescOk(fd) {
		return nil, storage.ErrInvalidFile
	}
	m := &memFile{}
	m.open = true
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.ramFiles.Add(fd.String(), m)
	var cacheFile *os.File
	if (fd.Type == storage.TypeJournal || fd.Type == storage.TypeManifest) && ms.opt.LocalCacheDir != "" {
		var err error
		cacheFile, err = os.Create(path.Join(ms.opt.LocalCacheDir, fd.String()))
		if err != nil {
			return nil, err
		}
	}
	return &memWriter{memFile: m, ms: ms, fd: fd, cacheFile: cacheFile}, nil
}

// 删除数据
func (ms *S3Storage) Remove(fd storage.FileDesc) error {
	if !storage.FileDescOk(fd) {
		return storage.ErrInvalidFile
	}
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.ramFiles.Remove(fd.String())
	os.Remove(path.Join(ms.opt.LocalCacheDir, fd.String()))
	err := ms.objStore.Remove(fd.String())
	if err != nil {
		return err
	}
	return nil
}

func (ms *S3Storage) Rename(oldfd, newfd storage.FileDesc) error {
	if !storage.FileDescOk(oldfd) || !storage.FileDescOk(newfd) {
		return storage.ErrInvalidFile
	}
	if oldfd == newfd {
		return nil
	}
	return nil
}

func (ms *S3Storage) Close() error {
	log.Println("storage Close")
	return nil
}

type memFile struct {
	bytes.Buffer
	open bool
}

type memReader struct {
	*bytes.Reader
	ms     *S3Storage
	m      *memFile
	fd     storage.FileDesc
	closed bool
}

func (mr *memReader) Close() error {
	mr.ms.mu.Lock()
	defer mr.ms.mu.Unlock()
	if mr.closed {
		return storage.ErrClosed
	}
	mr.m.open = false
	mr.ms.ramFiles.Remove(mr.fd.String())
	return nil
}

type memWriter struct {
	*memFile
	ms        *S3Storage
	fd        storage.FileDesc
	cacheFile *os.File
}

// 同步
func (mw *memWriter) Sync() error {
	if mw.fd.Type == storage.TypeManifest && mw.cacheFile != nil {
		return mw.cacheFile.Sync()
	}
	return nil
}

// 写
func (mw *memWriter) Write(p []byte) (n int, err error) {
	n, err = mw.memFile.Write(p)
	if mw.cacheFile != nil {
		_, err2 := mw.cacheFile.Write(p)
		if err2 != nil {
			return 0, err2
		}
	}
	return n, err
}

func (mw *memWriter) Close() error {
	if mw.cacheFile != nil {
		tErr := mw.cacheFile.Close()
		if tErr != nil {
			return tErr
		}
	}
	fname := mw.fd.String()
	err := mw.ms.objStore.PutBytes(fname, mw.memFile.Bytes())
	if err != nil {
		return err
	}
	mw.memFile.open = false
	if mw.cacheFile != nil {
		os.Remove(mw.cacheFile.Name())
	}
	return nil
}
