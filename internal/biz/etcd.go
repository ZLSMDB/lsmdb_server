package biz

import (
	"context"

	"github.com/ZLSMDB/lsmdb_server/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// EtcdUsecase is a Etcd usecase.
type EtcdUsecase struct { /*  */
	repo data.EtcdRepo
	log  *log.Helper
}

// NewEtcdUsecase new a Etcd usecase.
func NewEtcdUsecase(repo data.EtcdRepo, logger log.Logger) *EtcdUsecase {
	return &EtcdUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (s *EtcdUsecase) Put(key string, value string, opt any) error {
	return s.repo.Put(context.Background(), key, value, opt)
}

func (s *EtcdUsecase) PutKV(key string, value string) error {
	return s.repo.PutKV(context.Background(), key, value)
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

func (s *EtcdUsecase) GetWithPrefix(prefix string) ([]string, []string) {
	return s.repo.GetWithPrefix(context.Background(), prefix)
}

func (s *EtcdUsecase) Grant(ttl int64) (clientv3.LeaseID, error) {
	return s.repo.Grant(context.Background(), ttl)
}

func (s *EtcdUsecase) KeepAliveOnce(leaseID clientv3.LeaseID) error {
	return s.repo.KeepAliveOnce(context.Background(), leaseID)
}
