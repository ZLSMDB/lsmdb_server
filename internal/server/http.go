package server

import (
	etcdV1 "github.com/ZLSMDB/lsmdb_server/api/etcd/v1"
	lsmdbv1 "github.com/ZLSMDB/lsmdb_server/api/lsmdb/v1"
	registerv1 "github.com/ZLSMDB/lsmdb_server/api/register/v1"
	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/ZLSMDB/lsmdb_server/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, lsmdbs *service.LsmdbService, etcd *service.EtcdService, register *service.RegisterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	lsmdbv1.RegisterLsmdbHTTPServer(srv, lsmdbs)
	etcdV1.RegisterEtcdHTTPServer(srv, etcd)
	registerv1.RegisterRegisterHTTPServer(srv, register)
	return srv
}
