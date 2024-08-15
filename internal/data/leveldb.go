package data

import (
	"sync"

	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tsandl/skvdb/leveldb"
	"github.com/tsandl/skvdb/leveldb/iterator"
	"github.com/tsandl/skvdb/leveldb/opt"
	"github.com/tsandl/skvdb/leveldb/util"
)

type leveldbRepo struct {
	leveldb *leveldb.DB
	conf    *conf.Data
	log     *log.Helper
}

// NewLevelRepo .
func NewLevelDBRepo(conf *conf.Data, leveldb *leveldb.DB, logger log.Logger) DBRepo {
	return &leveldbRepo{
		conf:    conf,
		leveldb: leveldb,
		log:     log.NewHelper(logger),
	}
}

const (
	KiB = 1024
	MiB = KiB * 1024
	GiB = MiB * 1024
)

type LevelDB struct {
	DB *leveldb.DB
}

// 配置文件结构
type Config struct {
	Endpoint         string
	MasterIp         []string
	Ak               string
	Sk               string
	Region           string
	DataDir          string
	LocationCacheDir string
}

func (l *leveldbRepo) CreateBucket(bucketName string) error {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(l.conf.Oss.AccessKey, l.conf.Oss.SecretKey, ""),
		Endpoint:         aws.String(l.conf.Oss.Endpoint),
		Region:           aws.String(l.conf.Oss.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		l.log.Errorf("create session fail")
	}
	s3Client := awss3.New(newSession)
	// 检查存储桶是否已存在
	_, err = s3Client.HeadBucket(&awss3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err == nil {
		// 存储桶已存在，返回已创建信息
		l.log.Infof("Bucket '%s' already exists.\n", bucketName)
		return nil
	} else {
		// 存储桶不存在，创建存储桶
		_, err := s3Client.CreateBucket(&awss3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			l.log.Error(err)
			return nil
		}
		l.log.Infof("Bucket '%s' created successfully.\n", bucketName)
	}
	return nil
}

// 创建db
func (l *leveldbRepo) NewLevelDBCli(bucketName string) error {
	// var m sync.Mutex
	opt := new(opt.Options)
	opt.CompactionTableSize = int(l.conf.Leveldb.CompactionTableTize * MiB)
	opt.IteratorSamplingRate = int(l.conf.Leveldb.IteratorSamplingRate * MiB)
	opt.WriteBuffer = int(l.conf.Leveldb.WriteBuffer * MiB)
	opt.BlockSize = int(l.conf.Leveldb.BlockSize * MiB)
	opt.BlockCacheCapacity = int(l.conf.Leveldb.BlockCacheCapacity * MiB)
	opt.Compression = 1
	// opt.WriteL0SlowdownTrigger = int(l.conf.Leveldb.WriteBuffer * MiB)
	// opt.WriteL0PauseTrigger = int(l.conf.Leveldb.WriteBuffer * MiB)

	localPath := l.conf.Leveldb.DataDir + bucketName
	s3opt := OpenOption{
		Bucket:        bucketName,
		Path:          localPath,
		Endpoint:      l.conf.Oss.Endpoint,
		Ak:            l.conf.Oss.AccessKey,
		Sk:            l.conf.Oss.SecretKey,
		Region:        l.conf.Oss.Region,
		LocalCacheDir: l.conf.Leveldb.CacheDir,
	}
	if err := l.CreateBucket(bucketName); err != nil {
		return err
	}

	s3LruCache := 512 * MiB
	if l.conf.Leveldb.S3LruCache != 0 {
		s3LruCache = int(l.conf.Leveldb.S3LruCache * MiB)
	}

	storage, err := NewS3Storage(s3opt, s3LruCache)
	if err != nil {
		l.log.Errorf("create s3 storage failed: because of %v", err)
		return err
	}
	// add lock
	// m.Lock()
	dbInstance, err := leveldb.Open(storage, opt)
	// m.Unlock()
	if err != nil {
		l.log.Errorf("create db %s failed: %v", bucketName, err)
		return err
	}
	l.leveldb = dbInstance
	return nil
}

// 写入数据
func (l *leveldbRepo) Set(key string, value []byte) error {
	return l.leveldb.Put([]byte(key), []byte(value), nil)
}

func (l *leveldbRepo) BatchSet(keys []string, values [][]byte) error {

	for i := 0; i < len(keys); i++ {
		if err := l.leveldb.Put([]byte(keys[i]), values[i], nil); err != nil {
			return err // 如果发生错误，立即返回
		}
	}
	return nil
}

// 读取数据
func (l *leveldbRepo) Get(key string) ([]byte, error) {
	data, err := l.leveldb.Get([]byte(key), nil)
	if err != nil {
		l.log.Errorf("get key %s value fail, err %v", key, err)
		return data, err
	}
	return data, nil
}

// 删除数据
func (l *leveldbRepo) Del(key string) error {
	return l.leveldb.Delete([]byte(key), nil)
}

// 查看db状态
func (l *leveldbRepo) State(value string) (string, error) {
	if value == "" {
		value = "leveldb.stats"
	}
	if value == "type" {
		return "leveldb", nil
	}

	return l.leveldb.GetProperty(value)
}

// 根据prefix迭代数据，返回类型为map
func (l *leveldbRepo) Iterator(prefix string) (map[string]string, error) {
	data := make(map[string]string)
	var iter iterator.Iterator
	if prefix == "" {
		iter = l.leveldb.NewIterator(nil, nil)
		for ok := iter.Seek([]byte("")); ok; ok = iter.Next() {
			data[string(iter.Key())] = string(iter.Value()[:])
		}
	} else {
		iter = l.leveldb.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
		for iter.Next() {
			data[string(iter.Key())] = string(iter.Value()[:])
		}
	}
	iter.Release()
	return data, iter.Error()
}

// 根据prefix获取数据，返回类型为string数组
func (l *leveldbRepo) IteratorOnlyKey(prefix string) ([]string, error) {
	data := make([]string, 0)
	var iter iterator.Iterator
	if prefix == "" {
		iter = l.leveldb.NewIterator(nil, nil)
		for ok := iter.Seek([]byte("")); ok; ok = iter.Next() {
			data = append(data, string(iter.Key()))
		}
	} else {
		iter = l.leveldb.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
		for iter.Next() {
			data = append(data, string(iter.Key()))
		}
	}
	iter.Release()
	return data, iter.Error()
}

// 打开db 与 上面的NewDB类似，在应用中主要用NewDB
func (l *leveldbRepo) OpenDB(bucketName string) (*leveldb.DB, error) {
	var m sync.Mutex
	opt := new(opt.Options)
	opt.CompactionTableSize = int(l.conf.Leveldb.CompactionTableTize * MiB)
	opt.IteratorSamplingRate = int(l.conf.Leveldb.IteratorSamplingRate * MiB)
	opt.WriteBuffer = int(l.conf.Leveldb.WriteBuffer * MiB)
	opt.BlockSize = int(l.conf.Leveldb.BlockSize * MiB)
	opt.BlockCacheCapacity = int(l.conf.Leveldb.BlockCacheCapacity * MiB)
	localPath := l.conf.Leveldb.DataDir + bucketName
	s3opt := OpenOption{
		Bucket:        bucketName,
		Path:          localPath,
		Endpoint:      l.conf.Oss.Endpoint,
		Ak:            l.conf.Oss.AccessKey,
		Sk:            l.conf.Oss.SecretKey,
		Region:        l.conf.Oss.Region,
		LocalCacheDir: l.conf.Leveldb.CacheDir,
	}
	// to judge bucket is exist
	if err := l.CreateBucket(bucketName); err != nil {
		return nil, err
	}
	s3LruCache := 512 * MiB
	if l.conf.Leveldb.S3LruCache != 0 {
		s3LruCache = int(l.conf.Leveldb.S3LruCache * MiB)
	}
	storage, err := NewS3Storage(s3opt, s3LruCache)
	if err != nil {
		l.log.Errorf("create s3 storage failed: %v", err)
		return nil, err
	}
	// add lock
	m.Lock()
	dbInstance, err := leveldb.Open(storage, opt)
	m.Unlock()
	if err != nil {
		l.log.Errorf("create db %s failed: %v", bucketName, err)
		return nil, err
	}
	return dbInstance, nil
}

// 关闭 DB
func (l *leveldbRepo) CloseDB() error {
	var m sync.Mutex
	m.Lock()
	err := l.leveldb.Close()
	m.Unlock()
	if err != nil {
		l.log.Errorf("fail close db err: %v", err)
	} else {
		l.log.Info("success close db")
	}
	return err
}
