package rd_client

// Role 权限定义表
type Role struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`   // 自增 ID
	Name        string `gorm:"unique;not null;type:varchar(255)" json:"name"` // 唯一角色名称，限制为 VARCHAR(255)
	Description string `gorm:"type:text" json:"description"`                  // 描述，使用 TEXT 类型
}

func (c *Role) TableName() string {
	return "rapid.client.role"
}
