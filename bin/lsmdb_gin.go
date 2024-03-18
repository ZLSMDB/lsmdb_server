package v1

// import (
// 	context "context"
// 	"fmt"
// 	"net/http"

// 	"github.com/ZLSMDB/lsmdb_server/internal/service"
// 	"github.com/gin-gonic/gin"
// 	kgin "github.com/go-kratos/gin"
// 	"github.com/go-kratos/kratos/v2/errors"
// 	"github.com/go-kratos/kratos/v2/middleware"
// 	"github.com/go-kratos/kratos/v2/middleware/recovery"
// 	"github.com/go-kratos/kratos/v2/transport"
// )

// import (
// 	context "context"
// 	"net/http"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	ht "github.com/go-kratos/kratos/v2/transport/http"
// 	"golang.org/x/sync/semaphore"
// )

// func RegisterLsmdbGinServer(s *ht.Server, srv LsmdbHTTPServer) {
// 	r := s.Route("/")
// 	r.GET("/lsmdb/opendb", _Lsmdb_Get0_HTTP_Handlers(srv))
// 	go _Lsmdb_Get0_GIN_Handler(srv)
// }

// func _Lsmdb_Get0_HTTP_Handlers(srv LsmdbHTTPServer) func(ctx gin.Context) {
// 	return func(ctx gin.Context) {
// 		sem := semaphore.NewWeighted(50) // 限制为50个并发
// 		Key := ctx.Param("key")
// 		// 定义一个通道来接收结果
// 		resultChan := make(chan *GetReply, 1)

// 		// 控制并发数
// 		if err := sem.Acquire(context.Background(), 1); err != nil {
// 			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to acquire semaphore"})
// 			return
// 		}

// 		go func() {
// 			defer sem.Release(1)
// 			reply, err := srv.Get(ctx, &GetRequest{Key: Key})
// 			if err == nil {
// 				resultChan <- reply
// 			} else {
// 				resultChan <- nil
// 			}
// 		}()

// 		// 在另一个 Goroutine 中等待异步读取完成
// 		go func() {
// 			defer close(resultChan)
// 			<-resultChan // 等待结果
// 		}()

// 		// 从通道中获取结果
// 		result := <-resultChan
// 		if result == nil || len(result.Value) == 0 {
// 			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "not found"})
// 			ctx.Data(ctx.Writer.Status(), ctx.ContentType(), nil)
// 			return
// 		}

// 		ctx.Data(http.StatusOK, "application/octet-stream", result.Value)
// 	}

// }

// func _Lsmdb_Get0_GIN_Handler(srv LsmdbHTTPServer) {
// 	router := gin.Default()
// 	rt := ht.NewServer()
// 	rt.Addr = "0.0.0.0:7891"

// 	// 添加跨域中间件
// 	config := cors.DefaultConfig()
// 	config.AllowOrigins = []string{"*"} // 允许所有来源
// 	router.Use(cors.New(config))

// 	// 控制并发数的信号量
// 	sem := semaphore.NewWeighted(50) // 限制为100个并发

// 	router.GET("/get/*key", func(ctx *gin.Context) {
// 		Key := ctx.Param("key")
// 		// 定义一个通道来接收结果
// 		resultChan := make(chan *GetReply, 1)

// 		// 控制并发数
// 		if err := sem.Acquire(context.Background(), 1); err != nil {
// 			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to acquire semaphore"})
// 			return
// 		}

// 		go func() {
// 			defer sem.Release(1)
// 			reply, err := srv.Get(ctx, &GetRequest{Key: Key})
// 			if err == nil {
// 				resultChan <- reply
// 			} else {
// 				resultChan <- nil
// 			}
// 		}()

// 		// 在另一个 Goroutine 中等待异步读取完成
// 		go func() {
// 			defer close(resultChan)
// 			<-resultChan // 等待结果
// 		}()

// 		// 从通道中获取结果
// 		result := <-resultChan
// 		if result == nil || len(result.Value) == 0 {
// 			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "not found"})
// 			ctx.Data(ctx.Writer.Status(), ctx.ContentType(), nil)
// 			return
// 		}

// 		ctx.Data(http.StatusOK, "application/octet-stream", result.Value)
// 	})

// 	router.Run(":7891")
// }

// func _Lsmdb_Get0_HTTP_Handler(srv LsmdbHTTPServer) func(ctx ht.Context) error {
// 	return func(ctx ht.Context) error {
// 		ctx = gin.Context{Request: &ctx.Request,
//         Writer:  w.ResponseWriter,}

// 		var in GetRequest
// 		if err := ctx.BindQuery(&in); err != nil {
// 			return err
// 		}
// 		if err := ctx.BindVars(&in); err != nil {
// 			return err
// 		}
// 		ht.SetOperation(ctx, OperationLsmdbGet)
// 		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
// 			return srv.Get(ctx, req.(*GetRequest))
// 		})
// 		out, err := h(ctx, &in)
// 		reply, err = srv.Get(&gin.Context{}, &GetRequest{Key: "yesy"})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Data(http.StatusOK, "application/octet-stream", result.Value)
// 		return ctx.Result(200, reply.Value)
// 	}
// }

// func customMiddleware(handler middleware.Handler) middleware.Handler {
// 	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
// 		if tr, ok := transport.FromServerContext(ctx); ok {
// 			fmt.Println("operation:", tr.Operation())
// 		}
// 		reply, err = handler(ctx, req)
// 		return
// 	}
// }

// func GIN(lsmdbs *service.LsmdbService) *gin.Engine {

// 	router := gin.Default()
// 	// 使用kratos中间件
// 	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))

// 	router.GET("/getBytes/{*key}", func(ctx *gin.Context) {
// 		key := ctx.Param("key")
// 		fmt.Println(key)

// 		if key == "error" {
// 			// 返回kratos error
// 			kgin.Error(ctx, errors.Unauthorized("auth_error", "no authentication"))
// 		} else {
// 			reply, err := lsmdbs.Get(ctx, &v1.GetRequest{Key: key})
// 			if err != nil {
// 				ctx.IndentedJSON(http.StatusNotFound, gin.H{key: "not found"})
// 				return
// 			}
// 			ctx.Data(ctx.Writer.Status(), ctx.ContentType(), reply.Value)
// 		}
// 	})

// 	return router
// }
