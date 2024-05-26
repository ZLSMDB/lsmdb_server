package data

import (
	"errors"

	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-kratos/kratos/v2/log"
	rocksdb "github.com/tecbot/gorocksdb"
	"github.com/tsandl/skvdb/leveldb"
)

type rocksdbRepo struct {
	rocksdb *rocksdb.DB
	conf    *conf.Data
	log     *log.Helper

	db     *rocksdb.DB
	opts   *rocksdb.Options
	ro     *rocksdb.ReadOptions
	wo     *rocksdb.WriteOptions
	dbPath string // in conf
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
func NewRocksdbRepo(conf *conf.Data, rocksdb *rocksdb.DB, logger log.Logger) DBRepo {
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
	repo.opts = rocksdb.NewDefaultOptions()
	repo.opts.SetCreateIfMissing(true)
	repo.ro = rocksdb.NewDefaultReadOptions()
	repo.wo = rocksdb.NewDefaultWriteOptions()
	var err error
	repo.rocksdb, err = rocksdb.OpenDb(repo.opts, bucketName)
	return err

}

func (repo *rocksdbRepo) Set(key string, value []byte) error {
	return repo.rocksdb.Put(repo.wo, []byte(key), value)
}

func (repo *rocksdbRepo) Get(key string) ([]byte, error) {
	value, err := repo.rocksdb.Get(repo.ro, []byte(key))
	if err != nil {
		return nil, err
	}
	defer value.Free()
	return value.Data(), nil
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
	repo.rocksdb.Close()
	repo.opts.Destroy()
	repo.ro.Destroy()
	repo.wo.Destroy()
	return nil
}
