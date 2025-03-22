package rx

import "github.com/gin-gonic/gin"

// WrapHandler 包裹处理请求的回调, 会转换 context, 使其具有自定义的方法
func WrapHandler(handler func(c *Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{Context: c}
		handler(ctx)
	}
}
