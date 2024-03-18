package server

// import (
// 	"github.com/ZLSMDB/lsmdb_server/api/app"
// 	"github.com/ZLSMDB/lsmdb_server/internal/conf"
// 	"github.com/ZLSMDB/lsmdb_server/internal/service"

// 	"github.com/go-kratos/kratos/v2/middleware/recovery"
// 	"github.com/go-kratos/kratos/v2/transport/http"
// )

// // NewHTTPServer new an HTTP server.
// func NewHTTPGINServer(c *conf.Server, lsmdbs *service.LsmdbService) *http.Server {
// 	var opts = []http.ServerOption{
// 		http.Middleware(
// 			recovery.Recovery(),
// 			MiddlewareCors(), //support cross-origin
// 		),
// 	}
// 	if c.Httpgin.Network != "" {
// 		opts = append(opts, http.Network(c.Httpgin.Network))
// 	}
// 	if c.Httpgin.Addr != "" {
// 		opts = append(opts, http.Address(c.Httpgin.Addr))
// 	}
// 	if c.Httpgin.Timeout != nil {
// 		opts = append(opts, http.Timeout(c.Httpgin.Timeout.AsDuration()))
// 	}

// 	srv := http.NewServer(opts...)

// 	// 启用 Keep-Alive, long connect
// 	// srv.SetKeepAlivesEnabled(true)

// 	// lsmdbv1.RegisterLsmdbHTTPServer(srv, lsmdbs)
// 	// registerv1.RegisterRegisterHTTPServer(srv, register)
// 	srv.HandlePrefix("/", app.GIN(lsmdbs))
// 	srv.HandlePrefix("/", app.GINMore(lsmdbs))
// 	return srv
// }
