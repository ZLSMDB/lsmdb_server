package data

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type etcdRepo struct {
	data *Data
	log  *log.Helper
}

// EtcdRepo is a Etcd repo.
type EtcdRepo interface {
	Put(ctx context.Context, key, value string, opt any) error
	PutKV(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) []byte
	Delete(ctx context.Context, key string) error
	DeleteWithPrefix(ctx context.Context, prefix string) error
	GetWithPrefix(ctx context.Context, prefix string) ([]string, []string)
	Grant(ctx context.Context, ttl int64) (clientv3.LeaseID, error)
	KeepAliveOnce(ctx context.Context, leaseID clientv3.LeaseID) error
}

// NewEtcdRepo .
func NewEtcdRepo(data *Data, logger log.Logger) EtcdRepo {
	return &etcdRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (e *etcdRepo) Put(ctx context.Context, key, value string, opt any) error {
	_, err := e.data.etcd.Put(ctx, key, value, opt.(clientv3.OpOption))
	if err != nil {
		e.log.Errorf("etcd: put data err: %v", err)
	}
	return err
}

func (e *etcdRepo) PutKV(ctx context.Context, key, value string) error {
	_, err := e.data.etcd.Put(ctx, key, value)
	if err != nil {
		e.log.Errorf("etcd: put data err: %v", err)
	}
	return err
}

func (e *etcdRepo) Get(ctx context.Context, key string) []byte {
	resp, err := e.data.etcd.Get(ctx, key)
	if err != nil {
		e.log.Errorf("etcd: get value failed, err:%v", err)
		return nil
	}
	var value []byte
	for _, ev := range resp.Kvs {
		value = ev.Value
	}
	return value
}

func (e *etcdRepo) Delete(ctx context.Context, key string) error {
	_, err := e.data.etcd.Delete(ctx, key)
	if err != nil {
		e.log.Errorf("etcd: delete key failed, err:%v", err)
		return err
	}
	return err
}

func (e *etcdRepo) DeleteWithPrefix(ctx context.Context, prefix string) error {
	_, err := e.data.etcd.Delete(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		e.log.Errorf("etcd: delete key with prefix failed, err:%v", err)
		return err
	}
	return err
}

func (e *etcdRepo) GetWithPrefix(ctx context.Context, prefix string) ([]string, []string) {
	resp, err := e.data.etcd.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		e.log.Errorf("etcd: get key with prefix failed, err:%v", err)
	}
	var values []string
	var keys []string
	for _, ev := range resp.Kvs {
		values = append(values, string(ev.Value))
		keys = append(keys, string(ev.Key))
	}
	return keys, values
}

// 获取db所在的机器，
// TODO 转化成获取key所在的region
func (e *etcdRepo) GetDBAddr(dbName string) string {
	dbname := e.Get(context.Background(), dbName)
	if dbname == nil {
		// "/node/node" 是blockserver节点的前缀，不可修改
		_, values := e.GetWithPrefix(context.Background(), "/node/node")
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randNum := r.Intn(100) % len(values)
		return values[randNum]
	}
	return string(dbname)
}

// 获取租约ID
func (e *etcdRepo) Grant(ctx context.Context, ttl int64) (clientv3.LeaseID, error) {
	resp, err := e.data.etcd.Grant(ctx, ttl)
	return resp.ID, err
}

// 续租

func (e *etcdRepo) KeepAliveOnce(ctx context.Context, leaseID clientv3.LeaseID) error {
	_, err := e.data.etcd.KeepAliveOnce(ctx, leaseID)
	return err
}
