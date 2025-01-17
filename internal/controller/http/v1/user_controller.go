package v1

import (
	"demo/internal/entity"
	"demo/internal/pkg/db"
	"demo/pkg/r"
)

type LoginPayload struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8"`
}

func Login(c *r.Context) {
	var payload LoginPayload
	var users []entity.RapidClientUser

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.Fail(&r.R{
			Error: err.Error(),
		})
		return
	}

	result := db.PgRapid.Find(&users)

	if result.Error != nil {
		c.Fail(&r.R{
			Error: result.Error.Error(),
		})
		return
	}
	
	c.Ok(&r.R{
		Data: &users,
	})
}
