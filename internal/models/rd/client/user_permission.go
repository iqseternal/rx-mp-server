package rd_client

import "time"

// UserPermission 权限表
type UserPermission struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`                          // 自增主键 ID
	Name        string    `gorm:"unique;not null;type:varchar(255)" json:"name"`                        // 唯一权限名称，最大长度 255
	Description string    `gorm:"type:text" json:"description"`                                         // 权限描述，使用 TEXT 类型
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null;" json:"created_at"` // 自动创建时间，不能为空并且添加索引
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"` // 自动更新时间，不能为空并且添加索引
}

func (c *UserPermission) TableName() string {
	return "rapid.client.user_permission"
}
