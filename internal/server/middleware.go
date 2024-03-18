package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// MiddlewareCors 设置跨域请求头
func MiddlewareCors() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if ts, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := ts.(http.Transporter); ok {
					ht.ReplyHeader().Set("Access-Control-Allow-Origin", "*")
					ht.ReplyHeader().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH,DELETE")
					ht.ReplyHeader().Set("Access-Control-Allow-Credentials", "false")
					ht.ReplyHeader().Set("Access-Control-Allow-Headers", "Content-Type,"+
						"X-Requested-With,Access-Control-Allow-Credentials,User-Agent,Content-Length,Authorization")
				}
			}
			return handler(ctx, req)
		}
	}
}

// MiddlewareCors sets CORS headers
// func MiddlewareCors() middleware.Middleware {
// 	return func(handler middleware.Handler) middleware.Handler {
// 		return func(ctx context.Context, req interface{}) (interface{}, error) {
// 			if ts, ok := transport.FromServerContext(ctx); ok {
// 				if ht, ok := ts.(http.Transporter); ok {
// 					req, ok := req.(http.Request)
// 					if !ok {
// 						// If the request is not an HTTP request, just proceed without modifying headers
// 						return handler(ctx, req)
// 					}

// 					origin := req.Header.Get("Origin")
// 					if origin != "" {
// 						ht.ReplyHeader().Set("Access-Control-Allow-Origin", origin)
// 						ht.ReplyHeader().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH,DELETE")
// 						ht.ReplyHeader().Set("Access-Control-Allow-Headers", "Content-Type,"+
// 							"X-Requested-With,Access-Control-Allow-Credentials,User-Agent,Content-Length,Authorization")
// 						ht.ReplyHeader().Set("Access-Control-Allow-Credentials", "false")

// 						if req.Method == "OPTIONS" {
// 							ht.ReplyHeader().Set("Access-Control-Max-Age", "43200") // 12 hours in seconds
// 						}
// 					}
// 				}
// 			}
// 			return handler(ctx, req)
// 		}
// 	}
// }



// func DefaultConfig() Config {
// 	return Config{
// 		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
// 		AllowCredentials: false,
// 		MaxAge:           12 * time.Hour,
// 	}
// }