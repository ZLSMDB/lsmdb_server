package server

import (
	"fmt"

	"github.com/ZLSMDB/lsmdb_server/api/app"
	"github.com/ZLSMDB/lsmdb_server/internal/conf"
	"github.com/ZLSMDB/lsmdb_server/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type HttpGinServer *http.Server

// NewHTTPServer new an HTTP server.
func NewHTTPGINServer(c *conf.Server, lsmdbs *service.LsmdbService, register *service.RegisterService, logger log.Logger) HttpGinServer {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			MiddlewareCors(), //support cross-origin
		),
	}
	if c.Httpgin.Network != "" {
		opts = append(opts, http.Network(c.Httpgin.Network))
	}
	if c.Httpgin.Addr != "" {
		opts = append(opts, http.Address(c.Httpgin.Addr))
	}
	if c.Httpgin.Timeout != nil {
		opts = append(opts, http.Timeout(c.Httpgin.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)

	// 启用 Keep-Alive, long connect
	// srv.SetKeepAlivesEnabled(true)
	fmt.Println("httpgin.................................")
	// lsmdbv1.RegisterLsmdbHTTPServer(srv, lsmdbs)
	// registerv1.RegisterRegisterHTTPServer(srv, register)
	srv.HandlePrefix("/", app.GIN(lsmdbs))
	return srv
}
