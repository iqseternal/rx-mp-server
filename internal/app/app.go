package app

import (
	"github.com/gin-gonic/gin"
	"rx-mp/config"
	"rx-mp/internal/controller"
)

func Run() {
	gin.SetMode(gin.DebugMode)

	engine := gin.New()

	controller.InitRouter(engine)

	err := engine.Run(":" + config.Http.Port)

	if err != nil {
		panic("启动失败")
	}
}
