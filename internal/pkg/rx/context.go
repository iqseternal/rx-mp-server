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

// AbortWithOk 立即终止返回成功状态
func (c *Context) AbortWithOk(data interface{}) {
	c.Ok(data)
	c.Abort()
}

// OkWithCode 返回成功, 默认message
func (c *Context) OkWithCode(bizCode int, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    bizCode,
		Message: biz.Message(bizCode),
		Data:    data,
	})
}

// AbortWithOkCode 立即终止返回成功状态, 伴随自定义状态码
func (c *Context) AbortWithOkCode(bizCode int, data interface{}) {
	c.OkWithCode(bizCode, data)
	c.Abort()
}

// OkWithMessage 返回成功, 默认code
func (c *Context) OkWithMessage(message string, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.Success,
		Message: &message,
		Data:    data,
	})
}

// OkWithCodeMessage 立即终止返回成功状态
func (c *Context) OkWithCodeMessage(bizCode int, message string, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    bizCode,
		Message: &message,
		Data:    data,
	})
}

// AbortWithOkMessage 立即终止返回成功状态, 伴随自定义信息
func (c *Context) AbortWithOkMessage(message string, data interface{}) {
	c.OkWithMessage(message, data)
	c.Abort()
}

// Fail 返回失败, 默认code、message
func (c *Context) Fail(data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.Failure,
		Data:    data,
		Message: biz.Message(biz.Failure),
	})
}

// AbortWithFail 立即终止返回失败状态
func (c *Context) AbortWithFail(data interface{}) {
	c.Fail(data)
	c.Abort()
}

// FailWithCode 返回失败, 默认message
func (c *Context) FailWithCode(bizCode int, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    bizCode,
		Data:    data,
		Message: biz.Message(bizCode),
	})
}

// AbortWithFailCode 立即终止返回失败状态, 伴随自定义状态码
func (c *Context) AbortWithFailCode(bizCode int, data interface{}) {
	c.FailWithCode(bizCode, data)
	c.Abort()
}

// FailWithMessage 返回失败, 默认code
func (c *Context) FailWithMessage(message string, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.Failure,
		Data:    data,
		Message: &message,
	})
}

// AbortWithFailMessage 立即完成中断请求, 返回失败状态, 伴随 自定义信息
func (c *Context) AbortWithFailMessage(message string, data interface{}) {
	c.FailWithMessage(message, data)
	c.Abort()
}

// FailWithCodeMessage 返回失败状态, 并伴随状态码和自定义信息
func (c *Context) FailWithCodeMessage(bizCode int, message string, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    bizCode,
		Message: &message,
		Data:    data,
	})
}

// Finish 完成当前请求, 自定义http状态码和rx返回
func (c *Context) Finish(httpStatus int, r *R) {
	c.JSON(httpStatus, r)
}

// AbortFinish 立即完成当前请求
func (c *Context) AbortFinish(httpStatus int, r *R) {
	c.Finish(httpStatus, r)
	c.Abort()
}
