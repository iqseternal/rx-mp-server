package api

import (
	"demo/pkg/r"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterRootController(router *gin.Engine) {
	router.GET("/api", r.WrapHandler(Root))
	router.GET("/api/t", r.WrapHandler(T))
}

func Root(c *r.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://rapid.oupro.cn")
}

func T(c *r.Context) {
	c.Ok(&r.R{
		Data: time.Now().UnixMilli(),
	})
}
