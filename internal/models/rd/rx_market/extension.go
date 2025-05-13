package rdMarket

import (
	"time"
)

type ExtensionStatusField struct {
	IsDeleted bool `json:"is_deleted"`
}

// Extension 插件
type Extension struct {
	ExtensionId      int64       `json:"extension_id" gorm:"primaryKey;autoIncrement;comment:'自增主键'"`
	ExtensionGroupId int64       `json:"extension_group_id" gorm:"type:int;not null;"`
	ExtensionUuid    string      `json:"extension_uuid" gorm:"column:extension_uuid;default:gen_random_uuid();not null;comment:'唯一标识'"`
	ExtensionName    string      `json:"extension_name" gorm:"not null;comment:'分组名称'"`
	UseVersion       *int64      `json:"use_version" gorm:"type:int;"`
	ScriptHash       *string     `json:"script_hash" gorm:"type:text;"`
	Metadata         interface{} `json:"metadata" gorm:"type:jsonb;default:'{}';serializer:json;not null;"`
	Description      *string     `json:"description" gorm:"type:varchar;comment:'描述信息'"`
	Enabled          int64       `json:"enabled" gorm:"default:1;not null;comment:'启用状态'"`

	Status ExtensionStatusField `json:"status" gorm:"type:jsonb;default:'{}';serializer:json;not null;"`

	CreatorID *uint64 `json:"creator_id" gorm:"type:int;"`
	UpdaterID *uint64 `json:"updater_id" gorm:"type:int;"`

	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
}

func (c *Extension) TableName() string {
	return "rapid.rx_market.extension"
}
