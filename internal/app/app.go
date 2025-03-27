package app

import (
	"rx-mp/config"
	"rx-mp/internal/router"

	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(gin.DebugMode)

	engine := gin.New()

	router.InitRouter(engine)

	err := engine.Run(":" + config.Http.Port)

	if err != nil {
		panic("启动失败")
	}
}
