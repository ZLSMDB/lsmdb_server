package data

import (
	"bytes"
	"fmt"
	"io"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/tsandl/skvdb/leveldb/storage"
)

type S3Client struct {
	s3Store *s3.S3
	opt     OpenOption
}

/* 此部分由block server调用 */
// 获取S3客户端
func GetS3Client(opt OpenOption) (*S3Client, error) {
	creds := credentials.NewStaticCredentials(opt.Ak, opt.Sk, "")
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}
	config := &aws.Config{
		Region:           aws.String(opt.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      creds,
		Endpoint:         aws.String(opt.Endpoint),
	}
	client := s3.New(session.New(config))
	testBucket := &s3.HeadBucketInput{
		Bucket: aws.String(opt.Bucket),
	}
	_, err = client.HeadBucket(testBucket)
	if err != nil {
		return nil, err
	}
	opt.Path = path.Clean(opt.Path)
	return &S3Client{s3Store: client, opt: opt}, nil
}

// 写入数据
func (client *S3Client) PutBytes(key string, data []byte) error {
	_, err := client.s3Store.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(client.opt.Bucket),
		// Key:    aws.String(client.opt.Path + "/" + key),
		Key:  aws.String(key),
		Body: bytes.NewReader(data),
	})
	return err
}

// 获取数据
func (client *S3Client) GetBytes(key string) ([]byte, error) {
	rsps, err := client.s3Store.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(client.opt.Bucket),
		// Key:    aws.String(client.opt.Path + "/" + key),
		Key: aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	// TODO 换成io.copy
	data, err := io.ReadAll(rsps.Body)
	return data, err
}

// 删除数据
func (client *S3Client) Remove(key string) error {
	_, err := client.s3Store.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(client.opt.Bucket),
		Key:    aws.String(key),
		// Key:    aws.String(client.opt.Path + "/" + key),
	})
	if err != nil {
		return err
	}
	return nil
}

// List数据
func (client *S3Client) List() ([]storage.FileDesc, error) {
	files := []storage.FileDesc{}
	err := client.s3Store.ListObjectsPages(&s3.ListObjectsInput{
		Bucket:    aws.String(client.opt.Bucket),
		Prefix:    aws.String("/"), // TODO
		Delimiter: aws.String("/"),
	},
		func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
			for _, obj := range p.Contents {
				fullName := *obj.Key
				_, relName := path.Split(fullName)
				fd, pOK := fsParseName(relName)
				if pOK {
					files = append(files, fd)
				}
			}
			return true
		},
	)
	return files, err
}

func fsParseName(name string) (fd storage.FileDesc, ok bool) {
	var tail string
	_, err := fmt.Sscanf(name, "%d.%s", &fd.Num, &tail)
	if err == nil {
		switch tail {
		case "log":
			fd.Type = storage.TypeJournal
		case "ldb", "sst":
			fd.Type = storage.TypeTable
		case "tmp":
			fd.Type = storage.TypeTemp
		default:
			return
		}
		return fd, true
	}
	n, _ := fmt.Sscanf(name, "MANIFEST-%d%s", &fd.Num, &tail)
	if n == 1 {
		fd.Type = storage.TypeManifest
		return fd, true
	}
	return
}
