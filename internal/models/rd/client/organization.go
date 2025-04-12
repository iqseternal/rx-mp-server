package rdClient

import "time"

// Organization 组织表, 表示某个组织
type Organization struct {
	ID          uint      `json:"id"          gorm:"primaryKey;autoIncrement;not null"`                  // 自增 ID
	Name        string    `json:"name"        gorm:"type:varchar(255);unique;not null"`                  // 唯一名称，VARCHAR(255)
	Description string    `json:"description" gorm:"type:varchar(500)"`                                  // 描述，VARCHAR(500)
	CreatedAt   time.Time `json:"created_at"  gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null;"` // 自动创建时间，索引
	UpdatedAt   time.Time `json:"updated_at"  gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null;"` // 自动更新时间，索引
	IsDeleted   int       `json:"is_deleted"  gorm:"column:is_deleted;type:int2;default:0"`
}

func (c *Organization) TableName() string {
	return "rapid.client.organization"
}
