package rdclient

import "time"

// Group 用户组
type Group struct {
	GroupId     uint      `json:"group_id"     gorm:"primaryKey;autoIncrement;not null"`
	GroupName   string    `json:"group_name"   gorm:"type:varchar(255)"`
	CreatedAt   uint      `json:"created_at"   gorm:""`
	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt   uint      `json:"updated_at"   gorm:""`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
	ParentId    uint      `json:"parent_id"    gorm:""`
}

func (c *Group) TableName() string {
	return "rapid.client.group"
}
