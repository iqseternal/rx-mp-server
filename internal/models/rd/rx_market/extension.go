package rdMarket

import "time"

type ExtensionStatusField struct {
	IsDeleted bool `json:"is_deleted"`
}

// Extension 插件
type Extension struct {
	ExtensionId      int         `json:"extension_id" gorm:"primaryKey;autoIncrement;comment:'自增主键'"`
	ExtensionGroupId int         `json:"extension_group_id" gorm:"type:int;not null;"`
	ExtensionUuid    string      `json:"extension_uuid" gorm:"default:gen_random_uuid();not null;comment:'唯一标识'"`
	ExtensionName    string      `json:"extension_name" gorm:"not null;comment:'分组名称'"`
	UseVersion       *int        `json:"use_version" gorm:"type:int;"`
	ScriptHash       *string     `json:"script_hash" gorm:"type:text;"`
	Metadata         interface{} `json:"metadata" gorm:"type:jsonb;serializer:json"`

	Status ExtensionStatusField `json:"status" gorm:"type:jsonb;serializer:json"`

	CreatorID *uint `json:"creator_id" gorm:"type:int;"`
	UpdaterID *uint `json:"updater_id" gorm:"type:int;"`

	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
}

func (c *Extension) TableName() string {
	return "rapid.rx_market.extension"
}
