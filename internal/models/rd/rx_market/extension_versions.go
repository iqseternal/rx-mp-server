package rdMarket

import "time"

// ExtensionVersions 插件
type ExtensionVersions struct {
	ExtensionVersionId int    `json:"extension_version_id" gorm:"primaryKey;autoIncrement;comment:'自增主键'"`
	ExtensionId        int    `json:"extension_id" gorm:"type:int;not null;"`
	ScriptContent      string `json:"script_content"`
	Version            int    `json:"version"`
	Description        string `json:"description"`

	CreatorID *uint `json:"creator_id"   gorm:"type:int;"`
	UpdaterID *uint `json:"updater_id"   gorm:"type:int;"`

	CreatedTime time.Time `json:"created_time" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP;not null"`
	UpdatedTime time.Time `json:"updated_time" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP;not null"`
}

func (c *ExtensionVersions) TableName() string {
	return "rapid.rx_market.extension_versions"
}
