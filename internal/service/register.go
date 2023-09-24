package service

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	pb "github.com/ZLSMDB/lsmdb_server/api/register/v1"
	"github.com/ZLSMDB/lsmdb_server/internal/biz"
	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type RegisterService struct {
	pb.UnimplementedRegisterServer
	etcdUc *biz.EtcdUsecase
	conf   *conf.Bootstrap
	log    *log.Helper
}

func NewRegisterService(conf *conf.Bootstrap, etcdUc *biz.EtcdUsecase, logger log.Logger) *RegisterService {
	srv := &RegisterService{conf: conf, etcdUc: etcdUc, log: log.NewHelper(logger)}
	nodeIPAddr, err := getIpAddr(srv.conf.Server.Grpc.Addr)
	if err != nil {
		panic(err)
	}
	leaseID, err := srv.etcdUc.Grant(10) // 设置租约时间为10秒
	if err != nil {
		srv.log.Errorf("grant lease ID fail:", err)
		panic(err)
	}
	err = srv.etcdUc.Put(fmt.Sprintf("node/%s", srv.conf.Data.NodeName), nodeIPAddr, clientv3.WithLease(leaseID))
	if err != nil {
		srv.log.Errorf("Failed to register node:", err)
		panic(err)
	}
	srv.log.Info("Node %s registered with IP %s\n", srv.conf.Data.NodeName, nodeIPAddr)
	// 定时发送心跳
	go func() {
		for {
			err := srv.etcdUc.KeepAliveOnce(leaseID)
			if err != nil {
				srv.log.Errorf("Failed to send heartbeat:", err)
			}
			time.Sleep(5 * time.Second) // 每隔5秒发送一次心跳
		}
	}()
	// 定时发送心跳
	return srv
}

func (s *RegisterService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}

func getIpAddr(addr string) (string, error) {
	port := strings.Split(addr, ":")[1]
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	// TODO 一台机器多个ip情况得检测
	var ipAddr string
	for _, addr := range addrs {
		// 检查是否是IP地址，并且不是回环地址
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ipAddr = ipNet.IP.String()
				fmt.Println("IPv4 Address:", ipNet.IP.String())
			}
		}
	}
	addr = fmt.Sprintf("%s:%s", ipAddr, port)
	return addr, nil
}
