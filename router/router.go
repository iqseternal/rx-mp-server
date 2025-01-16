package router

import (
	"demo/handlers"
	"demo/libs/r"
	"demo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.LoggerToFile(), middleware.Cors())

	api := router.Group("/")
	//api.Use(middleware.JwtGuard())

	api.POST("/login", r.WrapHandler(handlers.Login))

	router.GET("/", r.WrapHandler(func(c *r.Context) {
		c.Ok(&r.R{})
	}))
}
