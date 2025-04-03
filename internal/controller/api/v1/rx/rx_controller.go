package v1_rx

import (
	"rx-mp/internal/middleware"
	"rx-mp/internal/pkg/rx"

	"github.com/gin-gonic/gin"
)

func RegisterRXController(router *gin.Engine) {

	routerGroup := router.Group("")
	routerGroup.Use(middleware.ResourceAccessControlMiddleware())
	{
		routerGroup.POST("/api/v1/rx/add_extension", rx.WrapHandler(AddExtension))
	}
}

func AddExtension(c *rx.Context) {

}
