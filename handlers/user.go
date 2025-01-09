package handlers

import (
	"demo/common"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {

	common.Logger.Info("这个事情不太妙啊")

	context.JSON(200, gin.H{
		"username": "john",
		"password": "secret",
	})
}
