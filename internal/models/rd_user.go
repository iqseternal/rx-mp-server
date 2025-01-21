package models

import "time"

type RdClientUser struct {
	ID       uint    `json:"id" gorm:"primaryKey;autoIncrement;not null"`     // 主键，自增，不能为空
	Username string  `json:"username" gorm:"type:varchar(255);not null;"`     // 用户名唯一，不能为空，最大长度为100
	Email    string  `json:"email" gorm:"type:varchar(255);unique;not null;"` // 邮箱唯一，不能为空，最大长度为255
	Password *string `json:"password" gorm:"type:varchar(300);"`              // 密码不能为空，最大长度为255（哈希值）
	Phone    *string `json:"phone" gorm:"type:varchar(20);"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"` // 创建时间，自动设置，不能为空
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"` // 更新时间，自动更新，不能为空
}

func (c *RdClientUser) TableName() string {
	return "rapid.client.user"
}

// RdClientOrganization 组织表, 表示某个组织
type RdClientOrganization struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`                          // 自增 ID
	Name        string    `gorm:"type:varchar(255);unique;not null" json:"name"`                        // 唯一名称，VARCHAR(255)
	Description string    `gorm:"type:varchar(500)" json:"description"`                                 // 描述，VARCHAR(500)
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null;" json:"created_at"` // 自动创建时间，索引
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"` // 自动更新时间，索引
}

func (c *RdClientOrganization) TableName() string {
	return "rapid.client.organization"
}

// RdClientUserOrganization 用户-组织关联表
type RdClientUserOrganization struct {
	ID             uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"` // 自增 ID，主键
	UserID         uint      `gorm:"not null;" json:"user_id"`                    // 用户 ID，不能为空，索引
	OrganizationID uint      `gorm:"not null;" json:"organization_id"`            // 组织 ID，不能为空，索引
	Role           string    `gorm:"not null;type:varchar(50)" json:"role"`       // 角色，不能为空，使用 VARCHAR(50)
	JoinedAt       time.Time `gorm:"not null;autoCreateTime" json:"joined_at"`    // 加入时间，不能为空，自动创建时间
}

func (c *RdClientUserOrganization) TableName() string {
	return "rapid.client.user_organization"
}

// RdClientRole 权限定义表
type RdClientRole struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`   // 自增 ID
	Name        string `gorm:"unique;not null;type:varchar(255)" json:"name"` // 唯一角色名称，限制为 VARCHAR(255)
	Description string `gorm:"type:text" json:"description"`                  // 描述，使用 TEXT 类型
}

func (c *RdClientRole) TableName() string {
	return "rapid.client.role"
}

// RdClientUserRole 角色表
type RdClientUserRole struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`                         // 自增主键 ID
	Name        string    `gorm:"unique;not null;type:varchar(255)" json:"name"`                       // 唯一名称，限制为 VARCHAR(255)
	RoleID      uint      `gorm:"not null;index" json:"role_id"`                                       // 角色 ID，不能为空，添加索引
	Description string    `gorm:"type:text" json:"description"`                                        // 描述，使用 TEXT 类型
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null" json:"created_at"` // 自动创建时间，不能为空
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null" json:"updated_at"` // 自动更新时间，不能为空
}

func (c *RdClientUserRole) TableName() string {
	return "rapid.client.user_role"
}

// RdClientPermission 权限表
type RdClientPermission struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`                          // 自增主键 ID
	Name        string    `gorm:"unique;not null;type:varchar(255)" json:"name"`                        // 唯一权限名称，最大长度 255
	Description string    `gorm:"type:text" json:"description"`                                         // 权限描述，使用 TEXT 类型
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null;" json:"created_at"` // 自动创建时间，不能为空并且添加索引
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"` // 自动更新时间，不能为空并且添加索引
}

func (c *RdClientPermission) TableName() string {
	return "rapid.client.user_permission"
}

// RdClientRolePermission 角色权限表
type RdClientRolePermission struct {
	RoleID       uint `gorm:"not null;" json:"role_id"`       // 角色 ID，普通字段
	PermissionID uint `gorm:"not null;" json:"permission_id"` // 权限 ID，普通字段
}

func (c *RdClientRolePermission) TableName() string {
	return "rapid.client.role_permission"
}
