package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// EtcdRepo is a Etcd repo.
type EtcdRepo interface {
	Put(ctx context.Context, key, value string, opt any) error
	Get(ctx context.Context, key string) []byte
	Delete(ctx context.Context, key string) error
	DeleteWithPrefix(ctx context.Context, prefix string) error
	GetWithPrefix(ctx context.Context, prefix string) []string
	Grant(ctx context.Context, ttl int64) (clientv3.LeaseID, error)
	KeepAliveOnce(ctx context.Context, leaseID clientv3.LeaseID) error
}

// EtcdUsecase is a Etcd usecase.
type EtcdUsecase struct { /*  */
	repo EtcdRepo
	log  *log.Helper
}

// NewEtcdUsecase new a Etcd usecase.
func NewEtcdUsecase(repo EtcdRepo, logger log.Logger) *EtcdUsecase {
	return &EtcdUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (s *EtcdUsecase) Put(key string, value string, opt any) error {
	return s.repo.Put(context.Background(), key, value, opt)
}

func (s *EtcdUsecase) Get(key string) string {
	return string(s.repo.Get(context.Background(), key))
}

func (s *EtcdUsecase) Delete(key string) error {
	return s.repo.Delete(context.Background(), key)
}

func (s *EtcdUsecase) DeleteWithPrefix(prefix string) error {
	return s.repo.DeleteWithPrefix(context.Background(), prefix)
}

func (s *EtcdUsecase) GetWithPrefix(prefix string) []string {
	return s.repo.GetWithPrefix(context.Background(), prefix)
}

func (s *EtcdUsecase) Grant(ttl int64) (clientv3.LeaseID, error) {
	return s.repo.Grant(context.Background(), ttl)
}

func (s *EtcdUsecase) KeepAliveOnce(leaseID clientv3.LeaseID) error {
	return s.repo.KeepAliveOnce(context.Background(), leaseID)
}
