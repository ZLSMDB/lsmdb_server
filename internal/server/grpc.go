package server

import (
	lsmdbv1 "github.com/ZLSMDB/lsmdb_server/api/lsmdb/v1"
	registerv1 "github.com/ZLSMDB/lsmdb_server/api/register/v1"
	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/ZLSMDB/lsmdb_server/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	g "google.golang.org/grpc"
)

const (
	MAX_MESSAGE_LENGTH = 256 * 1024 * 1024 // 可根据具体需求设置，此处设为64M
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, lsmdbs *service.LsmdbService, register *service.RegisterService, logger log.Logger) *grpc.Server {
	// MAX_MESSAGE_LENGTH := 256*1024*1024  // 可根据具体需求设置，此处设为256M
	// var opt = []g.ServerOption{
	// 	g.MaxRecvMsgSize(MAX_MESSAGE_LENGTH),
	// 	g.MaxSendMsgSize(MAX_MESSAGE_LENGTH),
	// }
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	// 设置grpc大小
	opts = append(opts, grpc.Options(g.MaxSendMsgSize(MAX_MESSAGE_LENGTH)))
	opts = append(opts, grpc.Options(g.MaxSendMsgSize(MAX_MESSAGE_LENGTH)))
	opts = append(opts, grpc.Options(g.MaxRecvMsgSize(MAX_MESSAGE_LENGTH)))
	// grpc.NewServer(grpc.StreamInterceptor(grpc.StreamInterceptor()),)
	// opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(size)))
	srv := grpc.NewServer(opts...)
	lsmdbv1.RegisterLsmdbServer(srv, lsmdbs)
	registerv1.RegisterRegisterServer(srv, register)
	return srv
}
