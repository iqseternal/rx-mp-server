package rd_client

import "time"

// UserOrganization 用户-组织关联表
type UserOrganization struct {
	ID             uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"` // 自增 ID，主键
	UserID         uint      `gorm:"not null;" json:"user_id"`                    // 用户 ID，不能为空，索引
	OrganizationID uint      `gorm:"not null;" json:"organization_id"`            // 组织 ID，不能为空，索引
	Role           string    `gorm:"not null;type:varchar(50)" json:"role"`       // 角色，不能为空，使用 VARCHAR(50)
	JoinedAt       time.Time `gorm:"not null;autoCreateTime" json:"joined_at"`    // 加入时间，不能为空，自动创建时间
}

func (c *UserOrganization) TableName() string {
	return "rapid.client.user_organization"
}
