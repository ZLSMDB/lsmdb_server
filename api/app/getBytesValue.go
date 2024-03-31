package app

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	v1 "github.com/ZLSMDB/lsmdb_server/api/lsmdb/v1"
	"github.com/ZLSMDB/lsmdb_server/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"golang.org/x/sync/semaphore"
)

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		// if tr, ok := transport.FromServerContext(ctx); ok {
		// 	fmt.Println("operation:", tr.Operation())
		// }
		reply, err = handler(ctx, req)
		return
	}
}

func GIN(lsmdbs *service.LsmdbService) *gin.Engine {
	// 禁止请求日志打印
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	// 添加跨域中间件
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"} // 允许所有来源
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(Cors())
	// 使用kratos中间件
	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))

	// 控制并发数的信号量
	sem := semaphore.NewWeighted(256) // 限制为100个并发

	router.GET("/gets/*key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		key = strings.TrimLeft(key, "/")
		// 定义一个通道来接收结果
		resultChan := make(chan *v1.GetReply, 1)

		// 控制并发数
		if err := sem.Acquire(context.Background(), 1); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to acquire semaphore"})
			return
		}

		// 异步处理
		go func() {
			defer sem.Release(1)
			reply, err := lsmdbs.Get(context.Background(), &v1.GetRequest{Key: key})
			if err != nil {
				resultChan <- nil
				return
			}
			resultChan <- reply
		}()

		// 等待异步处理结果
		result := <-resultChan
		close(resultChan)

		if result == nil || len(result.Value) == 0 {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{key: "not found"})
			return
		}

		ctx.Data(http.StatusOK, "application/octet-stream", result.Value)
	})

	router.GET("/get/*key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		key = strings.TrimLeft(key, "/")
		reply, err := lsmdbs.Get(ctx, &v1.GetRequest{Key: key})
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{key: "not found"})
			return
		}
		// ctx.Data(ctx.Writer.Status(), ctx.ContentType(), reply.Value)
		ctx.Data(ctx.Writer.Status(), "application/octet-stream", reply.Value)
	})

	return router
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func GINMores(lsmdbs *service.LsmdbService) *gin.Engine {

	router := gin.Default()

	// 添加跨域中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许所有来源
	router.Use(cors.New(config))

	// 控制并发数的信号量
	sem := semaphore.NewWeighted(100) // 限制为100个并发

	// router.GET("/get/:db/:key", func(ctx *gin.Context) {
	router.GET("/getMore/*key", func(ctx *gin.Context) {
		Key := ctx.Param("key")
		// 定义一个通道来接收结果
		resultChan := make(chan *v1.GetReply, 1)

		// 控制并发数
		if err := sem.Acquire(context.Background(), 1); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to acquire semaphore"})
			return
		}

		go func() {
			defer sem.Release(1)

			reply, err := lsmdbs.Get(context.Background(), &v1.GetRequest{Key: Key})
			if err == nil {
				resultChan <- reply
			} else {
				resultChan <- nil
			}
		}()

		// 在另一个 Goroutine 中等待异步读取完成
		go func() {
			defer close(resultChan)
			<-resultChan // 等待结果
		}()

		// 从通道中获取结果
		result := <-resultChan
		if result == nil || len(result.Value) == 0 {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "not found"})
			ctx.Data(ctx.Writer.Status(), ctx.ContentType(), nil)
			return
		}

		ctx.Data(http.StatusOK, "application/octet-stream", result.Value)
	})
	return router
}

func GINMore(lsmdbs *service.LsmdbService) *gin.Engine {
	router := gin.Default()

	// 添加跨域中间件
	router.Use(Cors())
	// 使用kratos中间件
	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))

	// 控制并发数的信号量
	sem := semaphore.NewWeighted(100) // 限制为100个并发

	router.GET("/fast/:key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		fmt.Println("key ", key)
		// 定义一个通道来接收结果
		resultChan := make(chan *v1.GetReply, 1)

		// 控制并发数
		if err := sem.Acquire(context.Background(), 1); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to acquire semaphore"})
			return
		}

		// 异步处理
		go func() {
			defer sem.Release(1)
			reply, err := lsmdbs.Get(context.Background(), &v1.GetRequest{Key: key})
			if err != nil {
				resultChan <- nil
				return
			}
			resultChan <- reply
		}()

		// 等待异步处理结果
		result := <-resultChan
		close(resultChan)

		if result == nil || len(result.Value) == 0 {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}

		ctx.Data(http.StatusOK, "application/octet-stream", result.Value)
	})

	return router
}
