package entity

type RapidClientUserGroup struct {
	Id int8 `json:"id"`
}

func (c *RapidClientUserGroup) TableName() string {
	return "rapid.client.user_group"
}
