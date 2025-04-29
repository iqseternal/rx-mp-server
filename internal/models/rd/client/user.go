package rdClient

import "time"

type User struct {
	UserID   uint64  `json:"user_id"  gorm:"primaryKey;autoIncrement;not null"`  // 主键，自增，不能为空
	Username string  `json:"username" gorm:"type:varchar(255);not null;"`        // 用户名唯一，不能为空，最大长度为100
	Email    string  `json:"email"    gorm:"type:varchar(255);unique;not null;"` // 邮箱唯一，不能为空，最大长度为255
	Password *string `json:"password" gorm:"type:varchar(300);"`                 // 密码不能为空，最大长度为255（哈希值）
	Phone    *string `json:"phone"    gorm:"type:varchar(20);"`

	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"` // 创建时间，自动设置，不能为空
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"` // 更新时间，自动更新，不能为空

	RefreshToken *string `json:"refresh_token,omitempty" gorm:"column:refresh_token"`
}

func (c *User) TableName() string {
	return "rapid.client.user"
}
