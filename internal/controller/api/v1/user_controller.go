package v1

import (
	"log"
	rd_client "rx-mp/internal/models/rd/client"
	"rx-mp/internal/pkg/common"
	"rx-mp/internal/pkg/jwt"
	"rx-mp/internal/pkg/storage"

	"rx-mp/internal/pkg/rx"

	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserController(router *gin.Engine) {
	router.POST("/api/v1/user/login", rx.WrapHandler(Login))
	router.POST("/api/v1/user/register", rx.WrapHandler(Register))
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

	var user rd_client.User
	result := storage.RdPostgress.
		Where("email = ?", payload.Email).
		Limit(1).
		First(&user)

	if result.Error != nil {
		c.FailWithMessage(result.Error.Error(), nil)
		return
	}

	accssToken, err := jwt.GenerateAccessToken(fmt.Sprint(user.UserID))
	if err != nil {
		fmt.Println("生成 access token 出错", err.Error())
		c.FailWithMessage(err.Error(), nil)
		return
	}

	refreshToken, err := jwt.GenerateRefershToken(fmt.Sprint(user.UserID))
	if err != nil {
		fmt.Println("生成 refresh token 出错", err.Error())
		c.FailWithMessage(err.Error(), nil)
		return
	}

	c.Ok(&rx.H{
		"tokens": &rx.H{
			"access_token":  accssToken,
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
	var user rd_client.User

	result := storage.RdPostgress.Where("email = ?", email).First(&user)

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
	user = rd_client.User{
		Email:    payload.Email,
		Username: payload.Username,
		Password: &password,
	}

	if user.Username == "" {
		// 生成随机hash
		user.Username = "用户" + common.GenerateRandomHexStr(5)
	}

	result = storage.RdPostgress.Create(&user)
	if result.Error != nil {
		c.FailWithMessage(result.Error.Error(), nil)
		return
	}

	c.Ok(&rx.R{
		Data: &user,
	})
}
