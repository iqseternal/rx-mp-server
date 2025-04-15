package api

import (
	"net/http"
	"rx-mp/internal/biz"
	rdMarket "rx-mp/internal/models/rd/rx_market"
	"rx-mp/internal/pkg/rx"
	"rx-mp/internal/pkg/storage"
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

	var extensionVersion rdMarket.ExtensionVersion
	result := storage.RdPostgres.First(&extensionVersion)

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	if result.RowsAffected == 0 {
		c.FailWithMessage("扩展组不存在", nil)
		return
	}

	c.Ok(time.Now().UnixMilli())
}
