package entity

type RapidClientUser struct {
	Id      int8 `json:"id"`
	GroupId int8 `json:"group_id"`
}

func (c *RapidClientUser) TableName() string {
	return "rapid.client.user"
}
