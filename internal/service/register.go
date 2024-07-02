package service

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
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
	conf   *conf.Bootstrap
	etcdUc *biz.EtcdUsecase
	log    *log.Helper
}

func NewRegisterService(conf *conf.Bootstrap, etcdUc *biz.EtcdUsecase, logger log.Logger) *RegisterService {
	srv := RegisterService{conf: conf, etcdUc: etcdUc, log: log.NewHelper(logger)}
	// nodeIPAddr, err := getIpAddr(srv.conf.Server.Grpc.Addr)
	nodeIPAddr, err := getIPPort(srv.conf.Server.Grpc.Addr, srv.conf.Data.NodeName)
	if err != nil {
		panic(err)
	}
	leaseID, err := srv.etcdUc.Grant(10) // 设置租约时间为10秒
	if err != nil {
		srv.log.Errorf("grant lease ID fail:", err)
		panic(err)
	} /*  */
	err = srv.etcdUc.Put(fmt.Sprintf("node/%s", srv.conf.Data.NodeName), nodeIPAddr, clientv3.WithLease(leaseID))
	if err != nil {
		srv.log.Errorf("Failed to register node:", err)
		panic(err)
	}
	srv.log.Infof("Node %s registered with IP %s\n", srv.conf.Data.NodeName, nodeIPAddr)
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
	return &srv
}

func (s *RegisterService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}

func readHosts(name string) (string, error) {
	file, err := os.Open("/etc/hosts")
	if err != nil {
		return "", fmt.Errorf("无法打开/etc/hosts文件: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		hostname := fields[1] // 请注意，这里假设主机名是在第二个字段
		ip := fields[0]

		if hostname == name {
			return ip, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("读取/etc/hosts文件时出现错误: %s", err)
	}

	return "", fmt.Errorf("未找到主机名为'%s'的IP地址", name)
}

func getIPPort(addr, nodeName string) (string, error) {
	port := strings.Split(addr, ":")[1]
	var ipAddr string
	ipAddr, err := readHosts(nodeName)
	if err != nil {
		return "", err
	}
	addr = fmt.Sprintf("%s:%s", ipAddr, port)
	return addr, nil
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
