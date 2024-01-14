package service

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/tsandl/skvdb/leveldb"
	"github.com/tsandl/skvdb/leveldb/opt"
	"github.com/tsandl/skvdb/leveldb/storage"
	eCli "go.etcd.io/etcd/client/v3"
	"google.golang.org/appengine/log"
)

type OpenOption struct {
	Bucket        string
	Path          string
	Ak            string
	Sk            string
	Region        string
	Endpoint      string
	LocalCacheDir string
}

func NewEtcdCli() *eCli.Client {
	cli, err := eCli.New(eCli.Config{
		Endpoints:   []string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"},
		DialTimeout: 5 * time.Second,
		// Username:    c.Etcd.UserName,
		// Password:    c.Etcd.Password,
	})
	if err != nil {
		panic(err)
	}
	return cli
}

func NewS3Cli(c *conf.Data) *s3.S3 {
	AccessKey, SecretKey := "minioadmin", "admin123456"
	Region, Endpoint := "cn-north-1", "10.176.34.182:9006"
	creds := credentials.NewStaticCredentials(AccessKey, SecretKey, "")
	_, err := creds.Get()
	if err != nil {
		panic(err)
	}
	config := &aws.Config{
		Region:           aws.String(Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      creds,
		Endpoint:         aws.String(Endpoint),
	}
	client := s3.New(session.New(config))
	return client
}

func OpenDB(bucketName string) (*leveldb.DB, error) {
	var m sync.Mutex
	opt := new(opt.Options)
	MiB := 1 * 1024 * 1024
	opt.CompactionTableSize = int(32 * MiB)
	opt.IteratorSamplingRate = int(4 * MiB)
	opt.WriteBuffer = int(32 * MiB)
	opt.BlockSize = int(32 * MiB)

	localPath := "/Users/tsan1024/Data/tmp/testing/" + bucketName
	s3opt := OpenOption{
		Bucket:        bucketName,
		Path:          localPath,
		Endpoint:      "",
		Ak:            "",
		Sk:            "",
		Region:        "",
		LocalCacheDir: "",
	}
	// to judge bucket is exist
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("", "", ""),
		Endpoint:         aws.String(""),
		Region:           aws.String(""),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession := session.New(s3Config)

	s3Client := awss3.New(newSession)
	bucket := aws.String(bucketName)
	cparams := &awss3.CreateBucketInput{
		Bucket: bucket, // Required
	}

	_, err := s3Client.CreateBucket(cparams)

	if err != nil {
		errMsg := err.Error()[0:23]
		if errMsg != "BucketAlreadyOwnedByYou" {
			log.Errorf(context.Background(), "Bucket already existed, bucket name: %s", errMsg)
			return nil, err
		}
	}
	storage, err := NewS3Storage(s3opt)
	if err != nil {
		log.Errorf(context.Background(), "create s3 storage failed: %v", err)
		return nil, err
	}
	// add lock
	m.Lock()
	dbInstance, err := leveldb.Open(storage, opt)
	m.Unlock()
	if err != nil {
		log.Errorf(context.Background(), "create db %s failed: %v", bucketName, err)
		return nil, err
	}
	return dbInstance, nil
}

func NewS3Storage(s3opt OpenOption) (storage.Storage, error) {
	fmt.Println(s3opt)
	panic("unimplemented")
}

func Mock() {

}

func TestOpenDB(t *testing.T) {

}
