package router

import (
	"demo/internal/controller/http/v1"
	"demo/internal/middleware"
	"demo/pkg/r"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.Cors())
	router.Use(gin.Logger(), gin.Recovery())

	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	api := router.Group("/")
	//api.Use(middleware.JwtGuard())

	api.POST("/login", r.WrapHandler(v1.Login))

	router.GET("/", r.WrapHandler(func(c *r.Context) {
		c.Ok(&r.R{
			Data: "不好",
		})
	}))
}
