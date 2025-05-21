package app

import (
	"github.com/gin-gonic/gin"
	"rx-mp/config"
	"rx-mp/internal/controller"
)

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

import _ "rx-mp/docs"

func Run() {
	gin.SetMode(gin.DebugMode)

	engine := gin.New()

	controller.InitRouter(engine)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := engine.Run(":" + config.Http.Port)

	if err != nil {
		panic("启动失败")
	}
}
