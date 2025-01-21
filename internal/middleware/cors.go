package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 处理前端跨域的中间件
func Cors() gin.HandlerFunc {

	config := cors.DefaultConfig()

	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AddAllowMethods("OPTIONS", "UPDATE")

	return cors.New(config)
}
