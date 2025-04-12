package rdClient

import "time"

// Role 权限定义表
type Role struct {
	RoleId      int       `json:"role_id"      gorm:"primaryKey;autoIncrement;not null"`
	RoleName    string    `json:"role_name"    gorm:"column:role_name;type:varchar;size:255"`
	CreatedAt   uint      `json:"created_at"   gorm:""`
	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt   uint      `json:"updated_at"   gorm:""`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
}

func (c *Role) TableName() string {
	return "rapid.client.role"
}
