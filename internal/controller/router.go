package controller

import (
	"fmt"
	"net/http"
	"rx-mp/internal/controller/api"
	"rx-mp/internal/controller/api/auth"
	"rx-mp/internal/controller/api/v1/rx"
	"rx-mp/internal/controller/api/v1/user"
	"rx-mp/internal/middleware"
	"rx-mp/internal/pkg/rx"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	SetupRouter(router)
	RegisterHandlers(router)
}

func SetupRouter(router *gin.Engine) {
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	router.NoRoute(rx.WrapHandler(noRoute))
	router.NoMethod(rx.WrapHandler(noMethod))
	router.LoadHTMLGlob("internal/templates/*")

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(middleware.DomainWhitelistMiddleware())
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.RecoveryMiddleware())
}

func RegisterHandlers(router *gin.Engine) {
	auth.RegisterAuthController(router)
	api.RegisterRootController(router)

	v1User.RegisterUserController(router)
	v1RX.RegisterRXController(router)
}

func noMethod(c *rx.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	c.FailWithMessage("Method not allowe", nil)
}

func noRoute(c *rx.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	c.FailWithMessage("Route not found", nil)
}
