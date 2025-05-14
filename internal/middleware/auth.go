package middleware

import (
	"fmt"
	"net/http"
	"rx-mp/internal/biz"
	"rx-mp/internal/models/rd/client"
	"rx-mp/internal/pkg/auth"
	"rx-mp/internal/pkg/common"
	"rx-mp/internal/pkg/mbic"
	"rx-mp/internal/pkg/rx"
	"rx-mp/internal/pkg/storage"
	"time"

	"github.com/gin-gonic/gin"
)

// ResourceAccessControlMiddleware 资源访问控制中间件, 会验证用户的权限
func ResourceAccessControlMiddleware() gin.HandlerFunc {

	return rx.WrapHandler(func(c *rx.Context) {
		accessAuthorization, err := c.Request.Cookie("access_token")

		if err != nil {
			c.FailWithCode(biz.NotCarryResourceAccessToken, nil)
			c.Abort()
			return
		}

		if accessAuthorization.Value == "dev_access_token" {
			var user *rdClient.User

			result := storage.RdPostgres.Where("user_id = ?", 10).First(&user)

			if result.Error != nil {
				c.FailWithCode(biz.UserNotExists, nil)
				c.Abort()
				return
			}

			mbic.SetMBICUser(c.Context, user)
			mbic.SetMBICUserID(c.Context, "10")
			c.Next()
			return
		}

		if err != nil {
			c.Finish(biz.Unauthorized, &rx.R{
				Code:    biz.Unauthorized,
				Message: biz.Message(biz.Unauthorized),
				Data:    nil,
			})
			c.Abort()
			return
		}

		accessToken := accessAuthorization.Value

		var user *rdClient.User
		userId, err := storage.MemoCache.Get(accessToken)
		if err != nil {
			claims, err := auth.VerifyAccessToken(accessToken)
			if err != nil {
				c.FailWithCode(biz.AccessTokenInvalid, nil)
				c.Abort()
				return
			}

			if time.Now().Unix() > claims.ExpiresAt.Unix() {
				c.FailWithCode(biz.AccessTokenExpired, nil)
				c.Abort()
				return
			}

			userId = claims.UserId
		}

		result := storage.RdPostgres.Where("user_id=?", userId).First(&user)
		if result.Error != nil {
			c.FailWithCode(biz.UserNotExists, nil)
			c.Abort()
			return
		}

		mbic.SetMBICUser(c.Context, user)
		mbic.SetMBICUserID(c.Context, userId)
		c.Next()
	})
}

// CredentialAccessControlMiddleware 凭证访问控制中间件, 验证 refresh token
func CredentialAccessControlMiddleware() gin.HandlerFunc {

	return rx.WrapHandler(func(c *rx.Context) {
		authorization := c.Request.Header.Get("Authorization")

		refreshToken, err := common.ParseBearerAuthorizationToken(authorization)
		if err != nil {
			c.Finish(biz.Unauthorized, &rx.R{
				Code:    biz.Unauthorized,
				Data:    nil,
				Message: biz.Message(biz.Unauthorized),
			})
			c.Abort()
			return
		}

		var user *rdClient.User
		result := storage.RdPostgres.Where("refresh_token=?", refreshToken).First(&user)
		if result.Error != nil {
			c.FailWithCode(biz.RefreshTokenInvalid, nil)
			c.Abort()
			return
		}

		// 验证 Refresh Token 有效性
		claims, err := auth.VerifyRefreshToken(refreshToken)
		if err != nil {
			c.Finish(http.StatusUnauthorized, &rx.R{
				Code:    biz.RefreshTokenInvalid,
				Message: biz.Message(biz.RefreshTokenInvalid),
				Data:    nil,
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			c.FailWithCode(biz.RefreshTokenExpired, nil)
			c.Abort()
			return
		}

		mbic.SetMBICUser(c.Context, user)
		mbic.SetMBICUserID(c.Context, fmt.Sprint(user.UserID))
		c.Next()
	})
}
