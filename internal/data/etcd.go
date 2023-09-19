package data

import (
	"context"
	"math/rand"
	"time"

	"github.com/ZLSMDB/lsmdb_server/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type etcdRepo struct {
	data *Data
	log  *log.Helper
}

// NewLevelRepo .
func NewEtcdRepo(data *Data, logger log.Logger) biz.EtcdRepo {
	return &etcdRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (e *etcdRepo) Put(ctx context.Context, key, value string) error {
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

func (e *etcdRepo) GetWithPrefix(ctx context.Context, prefix string) []string {
	resp, err := e.data.etcd.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		e.log.Errorf("etcd: get key with prefix failed, err:%v", err)
	}
	var value []string
	for _, ev := range resp.Kvs {
		value = append(value, string(ev.Value))
	}
	return value
}

// 获取db所在的机器，
// TODO 转化成获取key所在的region
func (e *etcdRepo) GetDBAddr(dbName string) string {
	dbname := e.Get(context.Background(), dbName)
	if dbname == nil {
		// "/node/node" 是blockserver节点的前缀，不可修改
		value := e.GetWithPrefix(context.Background(), "/node/node")
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randNum := r.Intn(100) % len(value)
		return value[randNum]
	}
	return string(dbname)
}
