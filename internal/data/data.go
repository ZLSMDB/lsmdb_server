package data

import (
	"time"

	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	eCli "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEtcdRepo, NewEtcdCli, NewRepo, NewS3Cli, NewOssRepo, NewRedisCli)

// Data .
type Data struct {
	// TODO wrapped database client
	osscli *s3.S3
	rdb    *redis.Client
	etcd   *eCli.Client
}

// NewData .
func NewData(etcd *eCli.Client, s3 *s3.S3, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")

	}
	return &Data{etcd: etcd, osscli: s3, rdb: rdb}, cleanup, nil
}

func NewEtcdCli(c *conf.Data) *eCli.Client {
	cli, err := eCli.New(eCli.Config{
		Endpoints:   c.Etcd.ClusterAddrs,
		DialTimeout: 5 * time.Second,
		// Username:    c.Etcd.UserName,
		// Password:    c.Etcd.Password,
	})
	if err != nil {
		log.Errorf("connect etcd fail %v", err)
		panic(err)
	}
	return cli
}

func NewS3Cli(c *conf.Data) *s3.S3 {
	creds := credentials.NewStaticCredentials(c.Oss.AccessKey, c.Oss.SecretKey, "")
	_, err := creds.Get()
	if err != nil {
		panic(err)
	}
	config := &aws.Config{
		Region:           aws.String(c.Oss.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      creds,
		Endpoint:         aws.String(c.Oss.Endpoint),
	}
	client := s3.New(session.New(config))
	return client
}

func NewRedisCli(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password, // no password set
		DB:       int(c.Redis.Db),  // use default DB
	})
	return rdb
}
