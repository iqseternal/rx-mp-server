package auth

import (
	"net/http"
	"rx-mp/internal/biz"
	"rx-mp/internal/pkg/jwt"
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
	refreshToken := r.GetHeader("Authorization")
	if refreshToken == "" {
		r.Finish(http.StatusUnauthorized, &rx.R{
			Code:    biz.BizRefreshTokenInvalid,
			Message: biz.BizMessage(biz.BizRefreshTokenInvalid),
			Data:    nil,
		})
		return
	}

	// 验证 Refresh Token 有效性
	claims, err := jwt.VerifyRefershToken(refreshToken)
	if err != nil {
		r.Finish(http.StatusUnauthorized, &rx.R{
			Code:    biz.BizRefreshTokenInvalid,
			Message: biz.BizMessage(biz.BizRefreshTokenInvalid),
			Data:    nil,
		})
		return
	}

	// 生成新 Token（复用 GetAccessToken 逻辑）
	newAccessToken, _ := jwt.GenerateAccessToken(claims.UserId)

	r.Finish(http.StatusOK, &rx.R{
		Code:    biz.BizSuccess,
		Message: biz.BizMessage(biz.BizSuccess),
		Data:    newAccessToken,
	})
}

func UpdateAccessToken(r *rx.Context) {
	// TODO: 需要携带旧的 AccessToken
	refreshToken := r.GetHeader("Authorization")
	if refreshToken == "" {
		r.Finish(http.StatusUnauthorized, &rx.R{
			Code:    biz.BizRefreshTokenInvalid,
			Message: biz.BizMessage(biz.BizRefreshTokenInvalid),
			Data:    nil,
		})
		return
	}

	// 验证 Refresh Token 有效性
	claims, err := jwt.VerifyRefershToken(refreshToken)
	if err != nil {
		r.Finish(http.StatusUnauthorized, &rx.R{
			Code:    biz.BizRefreshTokenInvalid,
			Message: biz.BizMessage(biz.BizRefreshTokenInvalid),
			Data:    nil,
		})
		return
	}

	// 生成新 Token（复用 GetAccessToken 逻辑）
	newAccessToken, _ := jwt.GenerateAccessToken(claims.UserId)

	r.Finish(http.StatusOK, &rx.R{
		Code:    biz.BizSuccess,
		Message: biz.BizMessage(biz.BizSuccess),
		Data:    newAccessToken,
	})
}

func LogoutAccessToken(r *rx.Context) {
	accessToken := r.GetHeader("Authorization")
	if accessToken == "" {
		r.Finish(http.StatusBadRequest, &rx.R{
			Code:    biz.BizAccessTokenInvalid,
			Message: biz.BizMessage(biz.BizAccessTokenInvalid),
			Data:    nil,
		})
		return
	}

	// TODO: 删除 accessToken 的有效性
	// if err := biz.AddToBlacklist(token); err != nil {
	// 	r.Finish(http.StatusInternalServerError, &rx.R{
	// 		Code:    biz.BizAccessTokenInvalid,
	// 		Message: biz.BizMessage(biz.BizAccessTokenInvalid),
	// 		Data:    nil,
	// 	})
	// 	return
	// }

	r.Finish(http.StatusOK, &rx.R{
		Code:    biz.BizSuccess,
		Message: biz.BizMessage(biz.BizSuccess),
		Data:    nil,
	})
}

func GetRefreshToken(r *rx.Context) {
	r.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.BizMehodNotAllowed,
		Message: biz.BizMessage(biz.BizMehodNotAllowed),
		Data:    nil,
	})
	r.Abort()
}

func UpdateRefreshToken(r *rx.Context) {
	r.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.BizMehodNotAllowed,
		Message: biz.BizMessage(biz.BizMehodNotAllowed),
		Data:    nil,
	})
	r.Abort()
}

func LogoutRefreshToken(r *rx.Context) {
	r.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.BizMehodNotAllowed,
		Message: biz.BizMessage(biz.BizMehodNotAllowed),
		Data:    nil,
	})
	r.Abort()
}
