package api

import (
	"net/http"
	"rx-mp/pkg/rx"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRootController(router *gin.Engine) {
	router.GET("/api", rx.WrapHandler(Root))
	router.GET("/api/t", rx.WrapHandler(T))
}

func Root(c *rx.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://rapid.oupro.cn")
}

func T(c *rx.Context) {

	c.Ok(&rx.R{
		Data: time.Now().UnixMilli(),
	})
}
