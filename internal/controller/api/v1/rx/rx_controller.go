package v1RX

import (
	"net/http"
	"rx-mp/internal/biz"
	"rx-mp/internal/middleware"
	"rx-mp/internal/pkg/rx"

	"github.com/gin-gonic/gin"
)

func RegisterRXController(router *gin.Engine) {

	routerGroup := router.Group("")
	routerGroup.Use(middleware.ResourceAccessControlMiddleware())
	{
		routerGroup.POST("/api/v1/rx/add_extension", rx.WrapHandler(AddExtension))
		routerGroup.POST("/api/v1/rx/remove_extension", rx.WrapHandler(RemoveExtension))
		routerGroup.POST("/api/v1/rx/active_extension", rx.WrapHandler(ActiveExtension))
		routerGroup.POST("/api/v1/rx/deactive_extension", rx.WrapHandler(DeactiveExtension))
	}
}

func AddExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}

func RemoveExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}

func ActiveExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}

func DeactiveExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}
