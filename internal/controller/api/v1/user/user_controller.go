package v1User

import (
	"log"
	"rx-mp/internal/biz"
	"rx-mp/internal/middleware"
	rdClient "rx-mp/internal/models/rd/client"
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

// Login 用户登录接口
// @title          用户管理系统 API
// @version        1.0
// @description    基于 Gin + Gorm 的用户管理接口文档
// @host      localhost:8080
// @BasePath  /api/v1
func Login(c *rx.Context) {
	var payload LoginPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var user rdClient.User
	result := storage.RdPostgres.
		Where("email = ?", payload.Email).
		Limit(1).
		First(&user)

	if result.Error != nil {
		c.FailWithCode(biz.UserNotExists, nil)
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
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	err = storage.MemoCache.Set(refreshToken, fmt.Sprint(user.UserID))
	if err != nil {
		c.FailWithCodeMessage(biz.MemoryCacheQueryError, result.Error.Error(), nil)
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

// Register 注册用户
// @Title 创建用户
// @Description 通过用户信息创建新账户
// @Router /users [post]
// @Param user body rdClient.User true "用户对象"
func Register(c *rx.Context) {
	var payload RegisterPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var user *rdClient.User
	result := storage.RdPostgres.Where("email = ?", payload.Email).First(&user)

	if result.RowsAffected != 0 {
		c.FailWithCodeMessage(biz.Failure, "邮箱已存在", nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("密码哈希处理失败:", err)
		c.FailWithCodeMessage(biz.InternalServerError, "无法生成hash串", nil)
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
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	c.Ok(user)
}

// GetUserInfo 获取用户信息
// @Summary 根据ID查询用户
// @Description 通过用户ID获取完整用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Router /users/{id} [get]
func GetUserInfo(c *rx.Context) {
	user, err := mbic.GetMBICUser(c.Context)
	if err != nil {
		c.FailWithCodeMessage(biz.MBICQueryError, err.Error(), nil)
		return
	}

	c.Ok(user)
}
