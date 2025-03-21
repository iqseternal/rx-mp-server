package app

import (
	"rx-mp/config"
	"rx-mp/internal/router"

	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	router.InitRouter(r)

	err := r.Run(":" + config.Http.Port)

	if err != nil {
		panic("启动失败")
	}
}
