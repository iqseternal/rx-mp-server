package app

import (
	"demo/internal/pkg/config"
	"demo/internal/pkg/db"
	"demo/internal/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func Run(config *config.Config) {
	db.Init(config)

	r := gin.New()

	router.InitRouter(r)

	err := r.Run(":" + config.Http.Port)

	if err != nil {
		fmt.Println("启动失败")
	}
}
