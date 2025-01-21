package v1

import (
	"demo/internal/models"
	"demo/internal/pkg/common"
	"demo/internal/pkg/db"
	"demo/pkg/r"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func RegisterUserController(router *gin.Engine) {
	router.POST("/api/v1/user/login", r.WrapHandler(Login))
	router.POST("/api/v1/user/register", r.WrapHandler(Register))
}

type LoginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func Login(c *r.Context) {
	var payload LoginPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.Fail(&r.R{
			Error: err.Error(),
		})
		return
	}

	var user models.RdClientUser
	result := db.RdPg.
		Where("email = ?", payload.Email).
		Order("created_at desc").
		First(&user)

	if result.Error != nil {
		c.Fail(&r.R{
			Error: result.Error.Error(),
		})
		return
	}

	c.Ok(&r.R{
		Data: &user,
	})
}

type RegisterPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"omitempty,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8"`
}

func Register(c *r.Context) {
	var payload RegisterPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.Fail(&r.R{
			Error: err.Error(),
		})
		return
	}

	email := payload.Email
	var user models.RdClientUser

	result := db.RdPg.Where("email = ?", email).First(&user)

	if result.Error == nil {
		c.Fail(&r.R{
			Error: "email is exist",
		})
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
	user = models.RdClientUser{
		Email:    payload.Email,
		Username: payload.Username,
		Password: &password,
	}

	if user.Username == "" {
		// 生成随机hash
		user.Username = "用户" + common.GenerateRandomHexStr(5)
	}

	result = db.RdPg.Create(&user)
	if result.Error != nil {
		c.Fail(&r.R{
			Error: result.Error.Error(),
		})
		return
	}

	c.Ok(&r.R{
		Data: &user,
	})
}
