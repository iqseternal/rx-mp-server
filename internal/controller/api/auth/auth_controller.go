package auth

import (
	"fmt"
	"net/http"
	"rx-mp/internal/biz"
	"rx-mp/internal/middleware"
	"rx-mp/internal/models/rd/client"
	"rx-mp/internal/pkg/auth"
	"rx-mp/internal/pkg/common"
	"rx-mp/internal/pkg/mbic"
	"rx-mp/internal/pkg/rx"
	"rx-mp/internal/pkg/storage"

	"github.com/gin-gonic/gin"
)

// RegisterAuthController 注册 auth 接口
func RegisterAuthController(router *gin.Engine) {
	tokenGroup := router.Group("")
	tokenGroup.Use(middleware.CredentialAccessControlMiddleware())
	{
		tokenGroup.POST("/api/auth/get_access_token", rx.WrapHandler(GetAccessToken))
		tokenGroup.POST("/api/auth/update_access_token", rx.WrapHandler(UpdateAccessToken))
		tokenGroup.POST("/api/auth/logout_access_token", rx.WrapHandler(LogoutAccessToken))
	}

	refreshGroup := router.Group("")
	refreshGroup.Use(middleware.CredentialAccessControlMiddleware())
	{
		refreshGroup.POST("/api/auth/get_refresh_token", rx.WrapHandler(GetRefreshToken))
		refreshGroup.POST("/api/auth/update_refresh_token", rx.WrapHandler(UpdateRefreshToken))
		refreshGroup.POST("/api/auth/logout_refresh_token", rx.WrapHandler(LogoutRefreshToken))
	}
}

// GetAccessToken 接口：通过 refreshToken 来生成一个可用的 accessToken, 但是同时也会导致其他的 accessToken 失效
func GetAccessToken(r *rx.Context) {
	user, err := mbic.GetMBICUser(r.Context)
	if err != nil {
		r.Finish(http.StatusUnauthorized, &rx.R{
			Code:    biz.Unauthorized,
			Message: biz.Message(biz.Unauthorized),
			Data:    nil,
		})
		return
	}

	// 生成新 Token（复用 GetAccessToken 逻辑）
	newAccessToken, _ := auth.GenerateAccessToken(fmt.Sprint(user.UserID))
	err = storage.MemoCache.Set(newAccessToken, fmt.Sprint(user.UserID))
	if err != nil {
		r.FailWithMessage(err.Error(), nil)
		return
	}

	r.Finish(http.StatusOK, &rx.R{
		Code:    biz.Success,
		Message: biz.Message(biz.Success),
		Data:    newAccessToken,
	})
}

// UpdateAccessToken 接口：更新 accessToken 并保证其他 accessToken 失效
func UpdateAccessToken(r *rx.Context) {
	// TODO: 需要携带旧的 refreshToken
	authorization := r.GetHeader("Authorization")

	refreshToken, err := common.ParseBearerAuthorizationToken(authorization)
	if err != nil {
		r.Finish(http.StatusUnauthorized, &rx.R{
			Code:    biz.BearerAuthorizationInvalid,
			Message: biz.Message(biz.BearerAuthorizationInvalid),
			Data:    nil,
		})
		return
	}

	var user *rdClient.User
	result := storage.RdPostgres.Where("refresh_token=?", refreshToken).First(&user)

	if result.Error != nil {
		r.Finish(http.StatusUnauthorized, &rx.R{
			Code:    biz.RefreshTokenInvalid,
			Message: biz.Message(biz.RefreshTokenInvalid),
			Data:    nil,
		})
		return
	}

	// 验证 Refresh Token 有效性
	claims, err := auth.VerifyRefreshToken(refreshToken)
	if err != nil {

		fmt.Println(err.Error())

		// 获取到了数据, 但是解析失败了?
		result := storage.RdPostgres.Model(&user).Where("refresh_token=?", refreshToken).Update("refresh_token", nil)

		if result.Error != nil {
			r.Finish(http.StatusInternalServerError, &rx.R{
				Code:    biz.DatabaseQueryError,
				Message: biz.Message(biz.DatabaseQueryError),
				Data:    nil,
			})
			return
		}

		r.Finish(http.StatusUnauthorized, &rx.R{
			Code:    biz.RefreshTokenInvalid,
			Message: biz.Message(biz.RefreshTokenInvalid),
			Data:    nil,
		})
		return
	}

	// 生成新 Token（复用 GetAccessToken 逻辑）
	newAccessToken, _ := auth.GenerateAccessToken(claims.UserId)
	err = storage.MemoCache.Set(newAccessToken, fmt.Sprint(user.UserID))
	if err != nil {
		r.FailWithMessage(err.Error(), nil)
		return
	}

	r.Finish(http.StatusOK, &rx.R{
		Code:    biz.Success,
		Message: biz.Message(biz.Success),
		Data:    newAccessToken,
	})
}

func LogoutAccessToken(r *rx.Context) {
	accessToken := r.GetHeader("Authorization")
	if accessToken == "" {
		r.Finish(http.StatusBadRequest, &rx.R{
			Code:    biz.AccessTokenInvalid,
			Message: biz.Message(biz.AccessTokenInvalid),
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
		Code:    biz.Success,
		Message: biz.Message(biz.Success),
		Data:    nil,
	})
}

func GetRefreshToken(r *rx.Context) {
	r.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.MethodNotAllowed,
		Message: biz.Message(biz.MethodNotAllowed),
		Data:    nil,
	})

	r.Abort()
}

func UpdateRefreshToken(r *rx.Context) {
	r.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.MethodNotAllowed,
		Message: biz.Message(biz.MethodNotAllowed),
		Data:    nil,
	})
	r.Abort()
}

func LogoutRefreshToken(r *rx.Context) {
	r.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.MethodNotAllowed,
		Message: biz.Message(biz.MethodNotAllowed),
		Data:    nil,
	})
	r.Abort()
}
