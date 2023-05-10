package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			//AllowAllOrigins:  true,
			AllowOrigins:  []string{"*"}, // 等同于允许所有域名 #AllowAllOrigins:  true
			AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:  []string{"*", "Authorization"},
			ExposeHeaders: []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
			//跨域请求的资源凭证，被指定可以选择cookie、HTTP签名header等
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	)
}
