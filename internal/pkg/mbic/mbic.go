package mbic

import (
	"fmt"
	rd_client "rx-mp/internal/models/rd/client"

	"github.com/gin-gonic/gin"
)

const (
	MBICUser   = "user"
	MBICUserID = "user_id"
)

func SetMBICUser(c *gin.Context, user *rd_client.User) {
	c.Set(MBICUser, user)
}

func GetMBICUser(c *gin.Context) (*rd_client.User, error) {
	iuser, ihasUser := c.Get(MBICUser)
	if !ihasUser {
		return nil, fmt.Errorf("Not Found user with MBInc")
	}
	user := iuser.(*rd_client.User)
	return user, nil
}

func SetMBICUserID(c *gin.Context, userId string) {
	c.Set(MBICUserID, userId)
}

func GetMBICUserID(c *gin.Context) (int, error) {
	iuserId, ihasUserId := c.Get(MBICUserID)
	if !ihasUserId {
		return -1, fmt.Errorf("Not Found user id with MBInc")
	}
	user := iuserId.(int)
	return user, nil
}
