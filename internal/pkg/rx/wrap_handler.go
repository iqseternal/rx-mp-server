package rx

import (
	"github.com/gin-gonic/gin"
)

// GetWrapRXContext 手动包裹 context, 获得符合 rx 的 context
func GetWrapRXContext(c *gin.Context) *Context {

	return &Context{
		Context: c,
	}
}

// WrapHandler 包裹处理请求的回调, 会转换 context, 使其具有自定义的方法
func WrapHandler(handler func(c *Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(GetWrapRXContext(c))
	}
}
