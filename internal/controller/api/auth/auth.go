package auth

import (
	"net/http"
	"rx-mp/internal/biz"
	"rx-mp/internal/pkg/rx"

	"github.com/gin-gonic/gin"
)

func RegisterAuthController(router *gin.Engine) {
	router.POST("/api/auth/get_access_token", rx.WrapHandler(GetAccessToken))
	router.POST("/api/auth/update_access_token", rx.WrapHandler(UpdateAccessToken))
	router.POST("/api/auth/logout_access_token", rx.WrapHandler(LogoutAccessToken))

	router.POST("/api/auth/get_refresh_token", rx.WrapHandler(GetRefreshToken))
	router.POST("/api/auth/update_refresh_token", rx.WrapHandler(UpdateRefreshToken))
	router.POST("/api/auth/logout_refresh_token", rx.WrapHandler(LogoutRefreshToken))
}

func GetAccessToken(r *rx.Context) {
	r.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.BizMehodNotAllowed,
		Message: biz.BizMessage(biz.BizMehodNotAllowed),
		Data:    nil,
	})
	r.Abort()
}

func GetRefreshToken(r *rx.Context) {

}

func UpdateAccessToken(r *rx.Context) {

}

func UpdateRefreshToken(r *rx.Context) {

}

func LogoutAccessToken(r *rx.Context) {

}

func LogoutRefreshToken(r *rx.Context) {

}
