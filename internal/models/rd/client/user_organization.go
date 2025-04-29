package rdClient

import "time"

// UserOrganization 用户-组织关联表
type UserOrganization struct {
	UserID         uint64    `json:"user_id"         gorm:"not null;"`               // 用户 ID，不能为空，索引
	OrganizationID uint64    `json:"organization_id" gorm:"not null;"`               // 组织 ID，不能为空，索引
	JoinedTime     time.Time `json:"joined_time"     gorm:"not null;autoCreateTime"` // 加入时间，不能为空，自动创建时间
	LeavedTime     time.Time `json:"leaved_time"     gorm:"column:leaved_time;type:timestamp"`
	IsLeaved       int64     `json:"is_leaved"       gorm:"column:is_leaved;type:int"`
}

func (c *UserOrganization) TableName() string {
	return "rapid.client.user_organization"
}
