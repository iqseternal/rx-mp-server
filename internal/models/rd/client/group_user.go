package rdClient

import "time"

// GroupUser 用户组与用户之间的关系
type GroupUser struct {
	Id          uint64    `json:"id"           gorm:"column:id;type:int8;primaryKey;autoIncrement;not null"`
	GroupId     uint64    `json:"group_id"     gorm:"primaryKey;not null"`
	UserId      uint64    `json:"user_id"      gorm:"column:user_id;type:int8;not null"`
	CreatedAt   uint64    `json:"created_at"   gorm:""`
	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt   uint64    `json:"updated_at"   gorm:""`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
}

func (c *GroupUser) TableName() string {
	return "rapid.client.group_user"
}
