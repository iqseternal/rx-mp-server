package router

import (
	"demo/handlers"
	"demo/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.LoggerToFile())

	router.Any("/", func(context *gin.Context) {
		date := time.Now().Format("2006/01/02 15:04:05")

		context.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"t": date,
			},
		})
	})

	router.POST("/login", handlers.Login)
}
