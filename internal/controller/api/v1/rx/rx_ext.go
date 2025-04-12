package v1RX

import (
	"net/http"
	"rx-mp/internal/biz"
	"rx-mp/internal/pkg/rx"
)

func AddExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}

func RemoveExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}

func ActiveExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}

func DeactiveExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}
