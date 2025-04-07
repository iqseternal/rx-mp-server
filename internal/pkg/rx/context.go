package rx

import (
	"net/http"
	"rx-mp/internal/biz"

	"github.com/gin-gonic/gin"
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
	Message *string `json:"message,omitempty"`
	// 具体的数据
	Data any `json:"data"`
	// 更多信息
	More *More `json:"more,omitempty"`
}

type Context struct {
	*gin.Context
}

// Ok 成功响应快捷方法, 默认code、message
func (c *Context) Ok(data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.Success,
		Message: biz.Message(biz.Success),
		Data:    data,
	})
}

// OkWithCode 返回成功, 默认message
func (c *Context) OkWithCode(bizCode int, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    bizCode,
		Message: biz.Message(bizCode),
		Data:    data,
	})
}

// OkWithMessage 返回成功, 默认code
func (c *Context) OkWithMessage(message string, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.Success,
		Message: &message,
		Data:    data,
	})
}

// Fail 返回失败, 默认code、message
func (c *Context) Fail(data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.Failure,
		Data:    data,
		Message: biz.Message(biz.Failure),
	})
}

// FailWithCode 返回失败, 默认message
func (c *Context) FailWithCode(bizCode int, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    bizCode,
		Data:    data,
		Message: biz.Message(bizCode),
	})
}

// FailWithMessage 返回失败, 默认code
func (c *Context) FailWithMessage(message string, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.Failure,
		Data:    data,
		Message: &message,
	})
}

// Finish 完成当前请求, 自定义http状态码和rx返回
func (c *Context) Finish(httpStatus int, r *R) {
	c.JSON(httpStatus, r)
}
