package auth

import (
	"demo/pkg/r"
	"github.com/gin-gonic/gin"
)

func RegisterAuthController(router *gin.Engine) {
	router.POST("/api/auth/get_access_token", r.WrapHandler(GetAccessToken))
	router.POST("/api/auth/get_refresh_token", r.WrapHandler(GetRefreshToken))

	router.POST("/api/auth/update_access_token", r.WrapHandler(UpdateAccessToken))
	router.POST("/api/auth/update_refresh_token", r.WrapHandler(UpdateRefreshToken))

	router.POST("/api/auth/logout_access_token", r.WrapHandler(LogoutAccessToken))
	router.POST("/api/auth/logout_refresh_token", r.WrapHandler(LogoutRefreshToken))
}

func GetAccessToken(r *r.Context) {

}

func GetRefreshToken(r *r.Context) {

}

func UpdateAccessToken(r *r.Context) {

}

func UpdateRefreshToken(r *r.Context) {

}

func LogoutAccessToken(r *r.Context) {

}

func LogoutRefreshToken(r *r.Context) {

}
