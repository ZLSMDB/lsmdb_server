package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// EtcdRepo is a Etcd repo.
type EtcdRepo interface {
	Put(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) []byte
	Delete(ctx context.Context, key string) error
	DeleteWithPrefix(ctx context.Context, prefix string) error
	GetWithPrefix(ctx context.Context, prefix string) []string
}

// EtcdUsecase is a Etcd usecase.
type EtcdUsecase struct {
	repo EtcdRepo
	log  *log.Helper
}

// NewEtcdUsecase new a Etcd usecase.
func NewEtcdUsecase(repo EtcdRepo, logger log.Logger) *EtcdUsecase {
	return &EtcdUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (s *EtcdUsecase) Put(key string, value string) error {
	return s.repo.Put(context.Background(), key, value)
}

func (s *EtcdUsecase) Get(key string) string {
	return string(s.repo.Get(context.Background(), key))
}

func (s *EtcdUsecase) Delete(key string) error {
	return s.repo.Delete(context.Background(), key)
}
