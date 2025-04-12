package v1User

import (
	"log"
	"rx-mp/internal/middleware"
	"rx-mp/internal/models/rd/client"
	"rx-mp/internal/pkg/auth"
	"rx-mp/internal/pkg/common"
	"rx-mp/internal/pkg/mbic"
	"rx-mp/internal/pkg/storage"

	"rx-mp/internal/pkg/rx"

	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserController(router *gin.Engine) {
	userPubGroup := router.Group("")
	{
		userPubGroup.POST("/api/v1/user/login", rx.WrapHandler(Login))
		userPubGroup.POST("/api/v1/user/register", rx.WrapHandler(Register))
	}

	userAuthGroup := router.Group("")
	userAuthGroup.Use(middleware.ResourceAccessControlMiddleware())
	{
		userAuthGroup.POST("/api/v1/user/get_user_info", rx.WrapHandler(GetUserInfo))
	}
}

type LoginPayload struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func Login(c *rx.Context) {
	var payload LoginPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithMessage(err.Error(), nil)
		return
	}

	var user rdClient.User
	result := storage.RdPostgres.
		Where("email = ?", payload.Email).
		Limit(1).
		First(&user)

	if result.Error != nil {
		c.FailWithMessage(result.Error.Error(), nil)
		return
	}

	accessToken, err := auth.GenerateAccessToken(fmt.Sprint(user.UserID))
	if err != nil {
		fmt.Println("生成 access token 出错", err.Error())
		c.FailWithMessage(err.Error(), nil)
		return
	}

	refreshToken, err := auth.GenerateRefreshToken(fmt.Sprint(user.UserID))
	if err != nil {
		fmt.Println("生成 refresh token 出错", err.Error())
		c.FailWithMessage(err.Error(), nil)
		return
	}

	result = storage.RdPostgres.
		Model(&user).
		Where("user_id=?", user.UserID).
		Update("refresh_token", refreshToken)

	if result.Error != nil {
		log.Println("更新 access token 出错:", result.Error)
		c.FailWithMessage(result.Error.Error(), nil)
		return
	}

	err = storage.MemoCache.Set(refreshToken, fmt.Sprint(user.UserID))
	if err != nil {
		c.FailWithMessage(result.Error.Error(), nil)
		return
	}

	c.Ok(&rx.H{
		"tokens": &rx.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}

type RegisterPayload struct {
	Email    string `json:"email"    binding:"required,email"`
	Username string `json:"username" binding:"omitempty,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8"`
}

func Register(c *rx.Context) {
	var payload RegisterPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithMessage(err.Error(), nil)
		return
	}

	email := payload.Email
	var user *rdClient.User

	result := storage.RdPostgres.Where("email = ?", email).First(user)

	if result.Error == nil {
		c.FailWithMessage("email is exist", nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("密码哈希处理失败:", err)
		c.JSON(500, gin.H{"error": "服务器内部错误"})
		return
	}

	password := string(hashedPassword)

	// 创建用户对象
	user = &rdClient.User{
		Email:    payload.Email,
		Username: payload.Username,
		Password: &password,
	}

	if user.Username == "" {
		// 生成随机hash
		user.Username = "用户" + common.GenerateRandomHexStr(5)
	}

	result = storage.RdPostgres.Create(&user)
	if result.Error != nil {
		c.FailWithMessage(result.Error.Error(), nil)
		return
	}

	c.Ok(user)
}

func GetUserInfo(c *rx.Context) {
	user, err := mbic.GetMBICUser(c.Context)
	if err != nil {
		c.FailWithMessage(err.Error(), nil)
		return
	}

	c.Ok(user)
}
