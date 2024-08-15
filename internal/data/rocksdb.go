package data

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-kratos/kratos/v2/log"
	gorocksdb "github.com/linxGnu/grocksdb"
	"github.com/tsandl/skvdb/leveldb"
)

type rocksdbRepo struct {
	rocksdb *gorocksdb.DB
	conf    *conf.Data
	log     *log.Helper
	opts    *gorocksdb.Options
	ro      *gorocksdb.ReadOptions
	wo      *gorocksdb.WriteOptions
	dbPath  string // in conf
}

func (l *rocksdbRepo) CreateBucket(bucketName string) error {
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

// NewLevelRepo .
func NewRocksdbRepo(conf *conf.Data, rocksdb *gorocksdb.DB, logger log.Logger) DBRepo {
	return &rocksdbRepo{
		conf:    conf,
		rocksdb: rocksdb,
		log:     log.NewHelper(logger),
	}
}
func (repo *rocksdbRepo) NewLevelDBCli(bucketName string) error {

	if err := repo.CreateBucket(bucketName); err != nil {
		return err
	}
	repo.opts = gorocksdb.NewDefaultOptions()
	repo.opts.SetCreateIfMissing(true)
	repo.ro = gorocksdb.NewDefaultReadOptions()
	repo.wo = gorocksdb.NewDefaultWriteOptions()
	var err error
	path := fmt.Sprintf("%s%s", repo.conf.Leveldb.DataDir, bucketName)
	repo.rocksdb, err = gorocksdb.OpenDb(repo.opts, path)
	return err

}

func (repo *rocksdbRepo) Set(key string, value []byte) error {
	return repo.rocksdb.Put(repo.wo, []byte(key), value)
}

func (repo *rocksdbRepo) BatchSets(keys []string, values [][]byte) error {
	for i := 0; i < len(keys); i++ {
		// err = repo.rocksdb.Put(repo.wo, []byte(keys[i]), values[i])
		if err := repo.rocksdb.Put(repo.wo, []byte(keys[i]), values[i]); err != nil {
			return err // 如果发生错误，立即返回
		}
	}
	return nil
}

func (repo *rocksdbRepo) BatchSet(keys []string, values [][]byte) error {

	writeBatch := gorocksdb.NewWriteBatch()
	defer writeBatch.Destroy() // 确保在函数结束时销毁 WriteBatch
	for i := 0; i < len(keys); i++ {
		writeBatch.Put([]byte(keys[i]), values[i])
	}

	// 创建写入选项
	writeOpts := gorocksdb.NewDefaultWriteOptions()
	defer writeOpts.Destroy() // 确保在函数结束时销毁写入选项

	// 批量写入
	if err := repo.rocksdb.Write(writeOpts, writeBatch); err != nil {
		log.Fatalf("Failed to write batch: %v", err)
	}

	repo.log.Info("Batch write successful!")
	return nil
}

// 对象池，用于重用内存
var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 1024) // 假设值的平均大小为1024字节
	},
}

func (repo *rocksdbRepo) Get(key string) ([]byte, error) {
	value, err := repo.rocksdb.Get(repo.ro, []byte(key))
	if err != nil {
		return nil, err
	}
	// defer value.Free()
	// // 将数据复制到字节切片中
	// data := make([]byte, len(value.Data()))
	// copy(data, value.Data())

	// 检查值是否为nil以避免解引用nil指针
	if value == nil {
		return nil, nil
	}

	// 获取值的数据
	data := value.Data()
	dataSize := len(data)

	// 从对象池获取缓冲区
	buf := bufPool.Get().([]byte)

	// 如果缓冲区不足以容纳数据，则分配新的缓冲区
	if cap(buf) < dataSize {
		buf = make([]byte, dataSize)
	} else {
		buf = buf[:dataSize]
	}

	// 复制数据到缓冲区
	copy(buf, data)

	// 释放原始值的内存
	value.Free()

	// 归还缓冲区到池中
	defer bufPool.Put(buf[:0])

	// 返回复制的数据
	return buf, nil
}

func (repo *rocksdbRepo) Del(key string) error {
	return repo.rocksdb.Delete(repo.wo, []byte(key))
}

func (repo *rocksdbRepo) State(value string) (string, error) {
	// Implement your custom logic here
	return "State function is not implemented yet", errors.New("not implemented")
}

func (repo *rocksdbRepo) Iterator(prefix string) (map[string]string, error) {
	result := make(map[string]string)
	iter := repo.rocksdb.NewIterator(repo.ro)
	defer iter.Close()

	for iter.Seek([]byte(prefix)); iter.ValidForPrefix([]byte(prefix)); iter.Next() {
		key := string(iter.Key().Data())
		value := string(iter.Value().Data())
		result[key] = value
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *rocksdbRepo) IteratorOnlyKey(prefix string) ([]string, error) {
	var keys []string
	iter := repo.rocksdb.NewIterator(repo.ro)
	defer iter.Close()

	for iter.Seek([]byte(prefix)); iter.ValidForPrefix([]byte(prefix)); iter.Next() {
		keys = append(keys, string(iter.Key().Data()))
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return keys, nil
}

func (repo *rocksdbRepo) OpenDB(bucketName string) (*leveldb.DB, error) {
	// nead to combine interface
	return nil, nil
}

func (repo *rocksdbRepo) CloseDB() error {
	repo.ro.Destroy()
	repo.wo.Destroy()
	repo.opts.Destroy()
	repo.rocksdb.Close()
	return nil
}
