package rdMarket

import "time"

// ExtensionVersion 插件版本记录, 某个插件始终会使用某个版本
type ExtensionVersion struct {
	ExtensionVersionId int64  `json:"extension_version_id" gorm:"primaryKey;autoIncrement;comment:'自增主键'"`
	ExtensionId        int64  `json:"extension_id" gorm:"type:int;not null;"`
	ScriptContent      string `json:"script_content"`
	Version            int64  `json:"version"`
	Description        string `json:"description"`

	CreatorID *uint64 `json:"creator_id"   gorm:"type:int;"`
	UpdaterID *uint64 `json:"updater_id"   gorm:"type:int;"`

	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
}

func (c *ExtensionVersion) TableName() string {
	return "rapid.rx_market.extension_version"
}
