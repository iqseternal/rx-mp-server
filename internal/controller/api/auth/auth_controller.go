package auth

import (
	"fmt"
	"net/http"
	"rx-mp/internal/biz"
	rd_client "rx-mp/internal/models/rd/client"
	"rx-mp/internal/pkg/common"
	"rx-mp/internal/pkg/jwt"
	"rx-mp/internal/pkg/rx"
	"rx-mp/internal/pkg/storage"

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

// GetAccessToken 接口：通过 refreshToken 来生成一个可用的 accessToken, 但是同时也会导致其他的 accessToken 失效
func GetAccessToken(r *rx.Context) {
	authorization := r.GetHeader("Authorization")

	refreshToken, err := common.ParseBearerAuthorizationToken(authorization)
	if err != nil {
		r.Finish(http.StatusUnauthorized, &rx.R{
			Code:    biz.BizBearerAuthorizationInvalid,
			Message: biz.BizMessage(biz.BizBearerAuthorizationInvalid),
			Data:    nil,
		})
		return
	}

	var user *rd_client.User
	result := storage.RdPostgress.Where("refresh_token=?", refreshToken).First(&user)

	if result.Error != nil {
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

		fmt.Println(err.Error())

		// 获取到了数据, 但是解析失败了?
		result := storage.RdPostgress.Model(&user).Where("refresh_token=?", refreshToken).Update("refresh_token", nil)

		if result.Error != nil {
			r.Finish(http.StatusInternalServerError, &rx.R{
				Code:    biz.BizDatabaseQueryError,
				Message: biz.BizMessage(biz.BizDatabaseQueryError),
				Data:    nil,
			})
			return
		}

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
