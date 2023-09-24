package service

import (
	"context"

	pb "github.com/ZLSMDB/lsmdb_server/api/etcd/v1"
	"github.com/ZLSMDB/lsmdb_server/internal/biz"
)

type EtcdService struct {
	pb.UnimplementedEtcdServer
	ucEtcd *biz.EtcdUsecase
}

func NewEtcdService(etcd *biz.EtcdUsecase) *EtcdService {
	return &EtcdService{ucEtcd: etcd}
}

func (s *EtcdService) Put(ctx context.Context, req *pb.EtcdPutRequest) (*pb.EtcdPutReply, error) {
	if err := s.ucEtcd.Put(req.Key, req.Value, nil); err != nil {
		return &pb.EtcdPutReply{Value: false}, err
	}
	return &pb.EtcdPutReply{Value: true}, nil
}

func (s *EtcdService) Get(ctx context.Context, req *pb.EtcdGetRequest) (*pb.EtcdGetReply, error) {
	return &pb.EtcdGetReply{Value: s.ucEtcd.Get(req.Key)}, nil
}
