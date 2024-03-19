package server

import (
	"github.com/ZLSMDB/lsmdb_server/api/app"
	lsmdbv1 "github.com/ZLSMDB/lsmdb_server/api/lsmdb/v1"
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
	// lsmdbv1.RegisterLsmdbHTTPServer(srv, lsmdbs)
	// registerv1.RegisterRegisterHTTPServer(srv, register)
	r := srv.Route("/")
	r.GET("/opendb/{dbname}", lsmdbv1.Lsmdb_OpenDBWeb0_HTTP_Handler(lsmdbs))
	r.GET("/close", lsmdbv1.Lsmdb_CloseDBWeb0_HTTP_Handler(lsmdbs))
	r.GET("/wget/{key:.*.*}", lsmdbv1.Lsmdb_Get0_HTTP_Handler(lsmdbs))
	srv.HandlePrefix("/", app.GIN(lsmdbs))
	return srv
}
