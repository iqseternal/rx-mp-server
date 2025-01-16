package r

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	Success = 0
	Failure = -1
)

type H = gin.H

type More struct {
	Pako bool `json:"pako"`
}

type R struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    H      `json:"data"`
	T       int64  `json:"t"`
	More    *More  `json:"more,omitempty"`
}

type Context struct {
	*gin.Context
}

func newSuccessResponse(context *gin.Context, r *R) {
	context.JSON(http.StatusOK, &R{
		Code:    Success,
		Message: r.Message,
		Data:    r.Data,
		T:       time.Now().UnixNano() / 1e6,
		More:    r.More,
	})
}

func newFailureResponse(context *gin.Context, r *R) {
	context.JSON(http.StatusOK, &R{
		Code:    Failure,
		Message: r.Message,
		Data:    r.Data,
		T:       time.Now().UnixNano() / 1e6,
		More:    r.More,
	})
}

func OK(context *gin.Context, r *R) {
	newSuccessResponse(context, r)
}

func Fail(context *gin.Context, r *R) {
	newFailureResponse(context, r)
}

func (c *Context) Ok(r *R) {
	OK(c.Context, r)
}

func (c *Context) Fail(r *R) {
	Fail(c.Context, r)
}

// WrapHandler 包裹处理请求的回调, 会转换 context, 使其具有自定义的方法
func WrapHandler(handler func(c *Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{Context: c}
		handler(ctx)
	}
}
