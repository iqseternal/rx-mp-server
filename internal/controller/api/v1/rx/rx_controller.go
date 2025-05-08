package v1RX

import (
	"rx-mp/internal/middleware"
	"rx-mp/internal/pkg/rx"

	"github.com/gin-gonic/gin"
)

func RegisterRXController(router *gin.Engine) {
	extensionRXOperatorGroup := router.Group("")
	extensionRXOperatorGroup.Use(middleware.ResourceAccessControlMiddleware())
	{
		extensionRXOperatorGroup.GET("/api/v1/rx/ext/get_ext_group_list", rx.WrapHandler(GetExtensionGroupList))

		extensionRXOperatorGroup.PUT("/api/v1/rx/ext/ext_group", rx.WrapHandler(AddExtensionGroup))
		extensionRXOperatorGroup.DELETE("/api/v1/rx/ext/ext_group", rx.WrapHandler(DelExtensionGroup))
		extensionRXOperatorGroup.GET("/api/v1/rx/ext/ext_group", rx.WrapHandler(GetExtensionGroup))
		extensionRXOperatorGroup.POST("/api/v1/rx/ext/ext_group", rx.WrapHandler(ModifyExtensionGroup))

		extensionRXOperatorGroup.GET("/api/v1/rx/ext/get_ext_list", rx.WrapHandler(GetExtensionList))

		extensionRXOperatorGroup.PUT("/api/v1/rx/ext/ext", rx.WrapHandler(AddExtension))
		extensionRXOperatorGroup.DELETE("/api/v1/rx/ext/ext", rx.WrapHandler(DelExtension))
		extensionRXOperatorGroup.GET("/api/v1/rx/ext/ext", rx.WrapHandler(GetExtension))
		extensionRXOperatorGroup.POST("/api/v1/rx/ext/ext", rx.WrapHandler(ModifyExtension))

		extensionRXOperatorGroup.GET("/api/v1/rx/ext/get_ext_version_list", rx.WrapHandler(GetExtensionVersionList))

		extensionRXOperatorGroup.PUT("/api/v1/rx/ext/ext/version", rx.WrapHandler(AddExtensionVersion))

		extensionRXOperatorGroup.POST("/api/v1/rx/ext/update_use_version", rx.WrapHandler(UpdateUseExtensionVersion))
	}

	extensionPublicGroup := router.Group("")
	extensionRXOperatorGroup.Use(middleware.ResourceAccessControlMiddleware())
	{
		extensionPublicGroup.GET("/api/v1/rx/ext/use", rx.WrapHandler(UseExtension))
		extensionPublicGroup.GET("/api/v1/rx/ext/use_group", rx.WrapHandler(UseExtensionGroup))
	}
}
