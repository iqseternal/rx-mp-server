package v1RX

import (
	"rx-mp/internal/middleware"
	"rx-mp/internal/pkg/rx"

	"github.com/gin-gonic/gin"
)

func RegisterRXController(router *gin.Engine) {

	routerGroup := router.Group("")
	routerGroup.Use(middleware.ResourceAccessControlMiddleware())
	{
		routerGroup.GET("/api/v1/rx/ext/get_ext_group_list", rx.WrapHandler(GetExtensionGroupList))

		routerGroup.PUT("/api/v1/rx/ext/ext_group", rx.WrapHandler(AddExtensionGroup))
		routerGroup.DELETE("/api/v1/rx/ext/ext_group", rx.WrapHandler(DelExtensionGroup))
		routerGroup.GET("/api/v1/rx/ext/ext_group", rx.WrapHandler(GetExtensionGroup))
		routerGroup.POST("/api/v1/rx/ext/ext_group", rx.WrapHandler(ModifyExtensionGroup))

		routerGroup.PUT("/api/v1/rx/ext/ext", rx.WrapHandler(AddExtension))
		routerGroup.DELETE("/api/v1/rx/ext/ext", rx.WrapHandler(DelExtension))
		routerGroup.GET("/api/v1/rx/ext/ext", rx.WrapHandler(GetExtension))
		routerGroup.POST("/api/v1/rx/ext/ext", rx.WrapHandler(ModifyExtension))
	}
}
