package rd_client

import "time"

// GroupUser
type GroupUser struct {
	Id          uint      `json:"id"           gorm:"column:id;type:int8;primaryKey;autoIncrement;not null"`
	GroupId     uint      `json:"group_id"     gorm:"primaryKey;not null"`
	UserId      uint      `json:"user_id"      gorm:"column:user_id;type:int8;not null"`
	CreatedAt   uint      `json:"created_at"   gorm:""`
	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt   uint      `json:"updated_at"   gorm:""`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
}

func (c *GroupUser) TableName() string {
	return "rapid.client.group_user"
}
