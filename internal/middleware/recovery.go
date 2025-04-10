package middleware

import (
	"github.com/gin-gonic/gin"
	"rx-mp/internal/pkg/rx"
)

// RecoveryMiddleware 错误处理中间件
func RecoveryMiddleware() gin.HandlerFunc {

	return rx.WrapHandler(func(c *rx.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.AbortWithFailMessage(c.Errors.Last().Error(), nil)
			return
		}
	})
}
