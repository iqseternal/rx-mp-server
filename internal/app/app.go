package app

import (
	"demo/internal/pkg/config"
	"demo/internal/router"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func Run() {
	r := gin.New()

	router.InitRouter(r)

	err := r.Run(":" + config.Http.Port)

	if err != nil {
		panic("启动失败")
	}
}
