package middleware

import (
	"net/http"
	"rx-mp/internal/biz"
	rd_client "rx-mp/internal/models/rd/client"
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
		authorization := c.Request.Header.Get("Authorization")

		accessToken, err := common.ParseBearerAuthorizationToken(authorization)
		if err != nil {
			c.Finish(biz.BizUnauthorized, &rx.R{
				Code: biz.BizUnauthorized,
				Data: nil,

				Message: biz.BizMessage(biz.BizUnauthorized),
			})
			c.Abort()
			return
		}

		var user *rd_client.User
		userId, err := storage.MemoCache.Get(accessToken)
		if err != nil {
			claims, err := auth.VerifyAccessToken(accessToken)
			if err != nil {
				c.FailWithCode(biz.BizBearerAuthorizationInvalid, nil)
				c.Abort()
				return
			}

			if time.Now().Unix() > claims.ExpiresAt.Unix() {
				c.FailWithCode(biz.BizBearerAuthorizationInvalid, nil)
				c.Abort()
				return
			}

			userId = claims.UserId
		}

		result := storage.RdPostgress.Where("user_id=?", userId).First(&user)
		if result.Error != nil {
			c.Finish(http.StatusUnauthorized, &rx.R{
				Code:    biz.BizUserNotExists,
				Message: biz.BizMessage(biz.BizUserNotExists),
				Data:    nil,
			})
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
			c.Finish(biz.BizUnauthorized, &rx.R{
				Code:    biz.BizUnauthorized,
				Data:    nil,
				Message: biz.BizMessage(biz.BizUnauthorized),
			})
			c.Abort()
			return
		}

		var user *rd_client.User
		result := storage.RdPostgress.Where("refresh_token=?", refreshToken).First(&user)
		if result.Error != nil {
			c.Finish(http.StatusUnauthorized, &rx.R{
				Code:    biz.BizBearerAuthorizationInvalid,
				Message: biz.BizMessage(biz.BizBearerAuthorizationInvalid),
				Data:    nil,
			})
			c.Abort()
			return
		}

		// 验证 Refresh Token 有效性
		claims, err := auth.VerifyRefershToken(refreshToken)
		if err != nil {
			c.Finish(http.StatusUnauthorized, &rx.R{
				Code:    biz.BizBearerAuthorizationInvalid,
				Message: biz.BizMessage(biz.BizBearerAuthorizationInvalid),
				Data:    nil,
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			c.FailWithCode(biz.BizBearerAuthorizationInvalid, nil)
			c.Abort()
			return
		}

		c.Set(mbic.MBICUser, user)
		c.Set(mbic.MBICUserID, user.UserID)
		c.Next()
	})
}
