package data

import (
	"fmt"
	"testing"

	"github.com/ZLSMDB/lsmdb_server/internal/biz"
	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func MockEtcd() *clientv3.Client {
	cli := NewEtcdCli(&conf.Data{Etcd: &conf.Data_Etcd{
		ClusterAddrs: []string{"127.0.0.1:2379"},
		UserName:     "",
		Password:     "",
	}})
	return cli
}

func MockEtcdRepo() biz.EtcdRepo {
	return NewEtcdRepo(&Data{etcd: MockEtcd()}, log.GetLogger())
}

func MockEtcdBiz() *biz.EtcdUsecase {
	return biz.NewEtcdUsecase(MockEtcdRepo(), log.GetLogger())
}
func TestConnect(t *testing.T) {
	etcdUc := MockEtcdBiz()
	err := etcdUc.Put("key", "value", nil)
	if err != nil {
		fmt.Println("put error")
	}
	fmt.Println(etcdUc.Get("key"))
	// fmt.Println(etcdUc.Get("key"))
}
