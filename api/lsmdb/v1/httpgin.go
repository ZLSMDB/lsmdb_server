package v1

import (
	context "context"

	http "github.com/go-kratos/kratos/v2/transport/http"
)

func Lsmdb_OpenDBWeb0_HTTP_Handler(srv LsmdbHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OpenDBWebRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLsmdbOpenDBWeb)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OpenDBWeb(ctx, req.(*OpenDBWebRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OpenDBWebReply)
		return ctx.Result(200, reply)
	}
}

func Lsmdb_CloseDBWeb0_HTTP_Handler(srv LsmdbHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CloseDBWebRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLsmdbCloseDBWeb)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CloseDBWeb(ctx, req.(*CloseDBWebRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CloseDBWebReply)
		return ctx.Result(200, reply)
	}
}

func Lsmdb_Get0_HTTP_Handler(srv LsmdbHTTPServer) func(ctx http.Context) error {
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
		// reply := out.(*GetReply)
		// return ctx.Blob(200, "application/octet-stream", reply.Value)
		return ctx.Blob(200, "application/octet-stream", out.(*GetReply).Value)
	}
}
