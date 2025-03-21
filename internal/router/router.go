package router

import (
	"fmt"
	"net/http"
	"rx-mp/internal/controller/api"
	"rx-mp/internal/controller/api/auth"
	v1 "rx-mp/internal/controller/api/v1"
	"rx-mp/internal/middleware"
	"rx-mp/pkg/rx"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.Cors())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.NoRoute(rx.WrapHandler(noRoute))
	router.NoMethod(rx.WrapHandler(noMethod))
	router.LoadHTMLGlob("internal/templates/*")

	auth.RegisterAuthController(router)
	api.RegisterRootController(router)

	v1.RegisterUserController(router)
}

func noMethod(c *rx.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	c.Fail(&rx.R{
		Error: "Method not allowe",
	})
}

func noRoute(c *rx.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	c.Fail(&rx.R{
		Error: "Route not found",
	})
}
