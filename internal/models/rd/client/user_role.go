package rd_client

import "time"

// UserRole 角色表
type UserRole struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`                         // 自增主键 ID
	Name        string    `gorm:"unique;not null;type:varchar(255)" json:"name"`                       // 唯一名称，限制为 VARCHAR(255)
	RoleID      uint      `gorm:"not null;index" json:"role_id"`                                       // 角色 ID，不能为空，添加索引
	Description string    `gorm:"type:text" json:"description"`                                        // 描述，使用 TEXT 类型
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null" json:"created_at"` // 自动创建时间，不能为空
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null" json:"updated_at"` // 自动更新时间，不能为空
}

func (c *UserRole) TableName() string {
	return "rapid.client.user_role"
}
