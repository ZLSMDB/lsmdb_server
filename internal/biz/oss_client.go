package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

// EtcdRepo is a Etcd repo.
type OssRepo interface {
	PutBytes(bucket string, key string, data []byte) error
	GetBytes(bucket string, key string) ([]byte, error)
}

// EtcdUsecase is a Etcd usecase.
type OssUsecase struct {
	repo OssRepo
	log  *log.Helper
}

// NewEtcdUsecase new a Etcd usecase.
func NewOssUsecase(repo OssRepo, logger log.Logger) *OssUsecase {
	return &OssUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (s *OssUsecase) PutBytes(bucket string, key string, data []byte) error {
	return s.repo.PutBytes(bucket, key, data)
}

func (s *OssUsecase) GetBytes(bucket string, key string) ([]byte, error) {
	return s.repo.GetBytes(bucket, key)
}
