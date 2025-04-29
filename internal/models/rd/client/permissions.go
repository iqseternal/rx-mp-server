package rdClient

import "time"

// Permissions 组织表, 表示某个组织
type Permissions struct {
	PermissionId int64     `json:"permission_id" gorm:"primaryKey;autoIncrement;not null"` // 自增 ID
	ResourceId   int64     `json:"resource_id"   gorm:"column:resource_id;type:int"`
	ResourceType int64     `json:"resource_type" gorm:"column:resource_type;type:int"`
	IsValid      int64     `json:"is_valid"      gorm:"column:is_valid;type:int"`
	CreatedAt    uint64    `json:"created_at"    gorm:""`
	CreatedTime  time.Time `json:"created_time"  gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt    uint64    `json:"updated_at"    gorm:""`
	UpdatedTime  time.Time `json:"updated_time"  gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
}

func (c *Permissions) TableName() string {
	return "rapid.client.permissions"
}
