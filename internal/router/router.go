package router

import (
	"demo/internal/controller/api"
	"demo/internal/controller/api/auth"
	v1 "demo/internal/controller/api/v1"
	"demo/internal/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	router.Use(middleware.Cors())
	router.Use(gin.Logger(), gin.Recovery())

	router.LoadHTMLGlob("internal/templates/*")
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	auth.RegisterAuthController(router)
	api.RegisterRootController(router)

	v1.RegisterUserController(router)
}
