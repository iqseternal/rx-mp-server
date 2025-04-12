package rdClient

import "time"

// Permissions 组织表, 表示某个组织
type Permissions struct {
	PermissionId int       `json:"permission_id" gorm:"primaryKey;autoIncrement;not null"` // 自增 ID
	ResourceId   int       `json:"resource_id"   gorm:"column:resource_id;type:int"`
	ResourceType int       `json:"resource_type" gorm:"column:resource_type;type:int"`
	IsValid      int       `json:"is_valid"      gorm:"column:is_valid;type:int"`
	CreatedAt    uint      `json:"created_at"    gorm:""`
	CreatedTime  time.Time `json:"created_time"  gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt    uint      `json:"updated_at"    gorm:""`
	UpdatedTime  time.Time `json:"updated_time"  gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
}

func (c *Permissions) TableName() string {
	return "rapid.client.permissions"
}
