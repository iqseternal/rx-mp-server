package rd_client

import "time"

// UserPermission 权限表
type UserPermission struct {
	ID          uint      `json:"id"          gorm:"primaryKey;autoIncrement;not null"`                  // 自增主键 ID
	Name        string    `json:"name"        gorm:"unique;not null;type:varchar(255)"`                  // 唯一权限名称，最大长度 255
	Description string    `json:"description" gorm:"type:text"`                                          // 权限描述，使用 TEXT 类型
	CreatedAt   time.Time `json:"created_at"  gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null;"` // 自动创建时间，不能为空并且添加索引
	UpdatedAt   time.Time `json:"updated_at"  gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null;"` // 自动更新时间，不能为空并且添加索引
}

func (c *UserPermission) TableName() string {
	return "rapid.client.user_permission"
}
