package rdClient

// UserRole 角色表
type UserRole struct {
	UserId uint64 `json:"user_id" gorm:"not null;index"` // 角色 ID，不能为空，添加索引
	RoleId uint64 `json:"role_id" gorm:"not null;index"` // 角色 ID，不能为空，添加索引
}

func (c *UserRole) TableName() string {
	return "rapid.client.user_role"
}
