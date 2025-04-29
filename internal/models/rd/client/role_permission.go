package rdClient

// RolePermission 角色权限表
type RolePermission struct {
	RoleID       uint64 `json:"role_id"       gorm:"not null;"` // 角色 ID，普通字段
	PermissionID uint64 `json:"permission_id" gorm:"not null;"` // 权限 ID，普通字段
}

func (c *RolePermission) TableName() string {
	return "rapid.client.role_permission"
}
