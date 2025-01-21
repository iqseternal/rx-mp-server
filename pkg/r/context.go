package r

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Success = 0
	Failure = -1
)

type H = gin.H

type More struct {
	// 数据是否经过压缩
	Pako bool `json:"pako"`
}

type R struct {
	// 成功返回, 一般成功为 Success, 失败为 Failure 或其他 http 码
	Code int `json:"code"`
	// 消息
	Error string `json:"error,omitempty"`
	// 具体的数据
	Data any `json:"data"`
	// 时间戳
	T int64 `json:"t"`
	// 更多信息
	More *More `json:"more,omitempty"`
}

type Context struct {
	*gin.Context
}

func newSuccessResponse(context *gin.Context, r *R) {
	context.JSON(http.StatusOK, &R{
		Code:  Success,
		Error: r.Error,
		Data:  r.Data,
		// 毫秒时间戳
		T:    time.Now().UnixMilli(),
		More: r.More,
	})
}

func newFailureResponse(context *gin.Context, r *R) {
	context.JSON(http.StatusOK, &R{
		Code:  Failure,
		Error: r.Error,
		Data:  r.Data,
		// 毫秒时间戳
		T:    time.Now().UnixMilli(),
		More: r.More,
	})
}

// Ok 带有特定包装的结构化成功返回
func Ok(context *gin.Context, r *R) {
	newSuccessResponse(context, r)
}

// Fail 带有特定包装的结构化失败返回
func Fail(context *gin.Context, r *R) {
	newFailureResponse(context, r)
}

// Ok 带有特定包装的结构化成功返回
func (c *Context) Ok(r *R) {
	Ok(c.Context, r)
}

// Fail 带有特定包装的结构化失败返回
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
