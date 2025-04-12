package mbic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rx-mp/internal/models/rd/client"
)

const (
	MBUser   = "user"
	MBUserID = "user_id"
)

func SetMBICUser(c *gin.Context, user *rdClient.User) {
	c.Set(MBUser, user)
}

func GetMBICUser(c *gin.Context) (*rdClient.User, error) {
	iUser, iHasUser := c.Get(MBUser)
	if !iHasUser {
		return nil, fmt.Errorf("not Found user with MBInc")
	}
	user := iUser.(*rdClient.User)
	return user, nil
}

func SetMBICUserID(c *gin.Context, userId string) {
	c.Set(MBUserID, userId)
}

func GetMBICUserID(c *gin.Context) (int, error) {
	iUserId, iHasUserId := c.Get(MBUserID)
	if !iHasUserId {
		return -1, fmt.Errorf("not Found user id with MBInc")
	}
	user := iUserId.(int)
	return user, nil
}
