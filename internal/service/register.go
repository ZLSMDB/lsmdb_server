package service

import (
	"context"

	pb "github.com/ZLSMDB/lsmdb_server/api/register/v1"
)

type RegisterService struct {
	pb.UnimplementedRegisterServer
}

func NewRegisterService() *RegisterService {
	return &RegisterService{}
}

func (s *RegisterService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
