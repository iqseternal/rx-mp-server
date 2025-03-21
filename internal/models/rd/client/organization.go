package rd_client

import "time"

// Organization 组织表, 表示某个组织
type Organization struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`                          // 自增 ID
	Name        string    `gorm:"type:varchar(255);unique;not null" json:"name"`                        // 唯一名称，VARCHAR(255)
	Description string    `gorm:"type:varchar(500)" json:"description"`                                 // 描述，VARCHAR(500)
	CreatedAt   time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null;" json:"created_at"` // 自动创建时间，索引
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null;" json:"updated_at"` // 自动更新时间，索引
}

func (c *Organization) TableName() string {
	return "rapid.client.organization"
}
