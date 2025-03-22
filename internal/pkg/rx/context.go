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

// Ok 成功响应快捷方法
func (c *Context) Ok(data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.BizSuccess,
		Message: biz.BizMessage(biz.BizSuccess),
		Data:    data,
	})
}

// OkWithCode
func (c *Context) OkWithCode(bizCode int, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    bizCode,
		Message: biz.BizMessage(bizCode),
		Data:    data,
	})
}

// OkWithMessage
func (c *Context) OkWithMessage(message string, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.BizSuccess,
		Message: &message,
		Data:    data,
	})
}

// Fail
func (c *Context) Fail(data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.BizFailure,
		Data:    data,
		Message: biz.BizMessage(biz.BizFailure),
	})
}

// FailWithCode
func (c *Context) FailWithCode(bizCode int, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    bizCode,
		Data:    data,
		Message: biz.BizMessage(bizCode),
	})
}

// FailWithMessage
func (c *Context) FailWithMessage(message string, data interface{}) {
	c.JSON(http.StatusOK, &R{
		Code:    biz.BizFailure,
		Data:    data,
		Message: &message,
	})
}

// Finish
func (c *Context) Finish(httpStatus int, r *R) {
	c.JSON(http.StatusOK, r)
}
