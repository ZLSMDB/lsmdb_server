package data

import (
	"bytes"
	"io"

	"github.com/ZLSMDB/lsmdb_server/internal/biz"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-kratos/kratos/v2/log"
)

type ossRepo struct {
	data *Data
	log  *log.Helper
}

// NewLevelRepo .
func NewOssRepo(data *Data, logger log.Logger) biz.OssRepo {
	return &ossRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 写入数据
func (client *ossRepo) PutBytes(bucket string, key string, data []byte) error {
	_, err := client.data.osscli.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})
	return err
}

// 获取数据
func (client *ossRepo) GetBytes(bucket string, key string) ([]byte, error) {
	rsps, err := client.data.osscli.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	// TODO 换成io.copy
	data, err := io.ReadAll(rsps.Body)
	return data, err
}
