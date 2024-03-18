package app

import (
	"context"
	"fmt"
	"net/http"

	v1 "github.com/ZLSMDB/lsmdb_server/api/lsmdb/v1"
	"github.com/ZLSMDB/lsmdb_server/internal/service"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
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

func GIN(lsmdbs *service.LsmdbService) *gin.Engine {

	router := gin.Default()
	// 使用kratos中间件
	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))

	router.GET("/getBytes/:name", func(ctx *gin.Context) {
		key := ctx.Param("name")
		fmt.Println(key)

		if key == "error" {
			// 返回kratos error
			kgin.Error(ctx, errors.Unauthorized("auth_error", "no authentication"))
		} else {
			reply, err := lsmdbs.Get(ctx, &v1.GetRequest{Key: key})
			if err != nil {
				ctx.IndentedJSON(http.StatusNotFound, gin.H{key: "not found"})
				return
			}
			ctx.Data(ctx.Writer.Status(), ctx.ContentType(), reply.Value)
		}
	})

	return router
}
