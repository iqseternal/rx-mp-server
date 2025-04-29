package rdMarket

import (
	"time"
)

type ExtensionGroupStatusField struct {
	IsDeleted bool `json:"is_deleted"`
}

// ExtensionGroup 插件
type ExtensionGroup struct {
	ExtensionGroupId   int64  `json:"extension_group_id" gorm:"primaryKey;autoIncrement;comment:'自增主键'"`
	ExtensionGroupUuid string `json:"extension_group_uuid" gorm:"default:gen_random_uuid();not null;comment:'唯一标识'"`
	ExtensionGroupName string `json:"extension_group_name" gorm:"not null;comment:'分组名称'"`
	Description        string `json:"description" gorm:"type:varchar;comment:'描述信息'"`
	Enabled            int64  `json:"enabled" gorm:"default:1;not null;comment:'启用状态'"`

	Status ExtensionGroupStatusField `json:"status" gorm:"type:jsonb;default:'{}';serializer:json;not null;comment:'状态信息'"`

	CreatorID *uint64 `json:"creator_id" gorm:"type:int;"`
	UpdaterID *uint64 `json:"updater_id" gorm:"type:int;"`

	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
}

func (c *ExtensionGroup) TableName() string {
	return "rapid.rx_market.extension_group"
}
