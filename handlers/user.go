package handlers

import (
	"demo/libs/r"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *r.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.Fail(&r.R{
			Message: "序列化错误",
		})
		return
	}

	c.Ok(&r.R{
		Data: r.H{
			"username": "HelloWorld",
		},
	})
}
