package v1

import (
	context "context"
	"fmt"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	http "github.com/go-kratos/kratos/v2/transport/http"
)

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}

func RegisterLsmdbHTTPServers(s *http.Server, srv LsmdbHTTPServer) {
	r := s.Route("/")
	r.GET("/get/{key:.*.*}", _Lsmdb_Get0_HTTP_Handlers(srv))
}

func _Lsmdb_Get0_HTTP_Handlers(srv LsmdbHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLsmdbGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*GetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetReply)
		return ctx.Blob(200, "application/octet-stream", reply.Value) // 有效
	}
}

func GetValue(srv LsmdbHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLsmdbGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*GetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetReply)
		return ctx.Result(200, reply.Value)
	}
}
