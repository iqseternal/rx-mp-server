package rdMarket

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// ExtensionView 插件版本视图
type ExtensionView struct {
	// 关联扩展插件
	ExtensionID   int64       `gorm:"column:extension_id;index:idx_extension"`
	ExtensionUUID string      `gorm:"column:extension_uuid;type:uuid;not null"`
	ExtensionName string      `gorm:"column:extension_name;index;type:varchar(255)"`
	UseVersion    *int64      `gorm:"column:use_version" binding:"required,gt=0"`
	ScriptHash    *string     `gorm:"column:script_hash" binding:"required"`
	Metadata      interface{} `gorm:"column:metadata;type:jsonb;serializer:json"`

	ExtensionCreatedTime     time.Time `gorm:"column:extension_created_time;->"`
	ExtensionUpdatedTime     time.Time `gorm:"column:extension_updated_time;->"`
	ExtensionCreatorID       *uint64   `gorm:"column:extension_creator_id"`
	ExtensionCreatorUsername string    `gorm:"->;column:extension_creator_username;type:varchar(64)"`
	ExtensionUpdaterUsername string    `gorm:"->;column:extension_updater_username;type:varchar(64)"`
	ExtensionUpdaterID       *uint64   `gorm:"column:extension_updater_id"`
	ExtensionDescription     *string   `gorm:"column:extension_description;type:varchar(500)"`

	// 关联扩展组
	ExtensionGroupID   int64  `gorm:"column:extension_group_id"`
	ExtensionGroupUUID string `gorm:"column:extension_group_uuid;type:uuid"`
	ExtensionGroupName string `gorm:"column:extension_group_name;type:varchar(255)"`

	// 基础版本字段
	ExtensionVersionId              int64     `gorm:"primaryKey;column:extension_version_id;autoIncrement:false"`
	ScriptContent                   string    `gorm:"column:script_content;type:text;"`
	Version                         int64     `gorm:"column:version;index:idx_version;comment:'版本号排序使用'"`
	ExtensionVersionCreatedTime     time.Time `gorm:"column:extension_version_created_time;->"`
	ExtensionVersionUpdatedTime     time.Time `gorm:"column:extension_version_updated_time;->"`
	ExtensionVersionCreatorID       *uint64   `gorm:"column:extension_version_creator_id"`
	ExtensionVersionCreatorUsername string    `gorm:"->;column:extension_version_creator_username;type:varchar(64)"`
	ExtensionVersionUpdaterUsername string    `gorm:"->;column:extension_version_updater_username;type:varchar(64)"`
	ExtensionVersionUpdaterID       *uint64   `gorm:"column:extension_version_updater_id"`
	ExtensionVersionDescription     *string   `gorm:"column:extension_version_description;type:varchar(500)"`

	IsEnabled int `gorm:"column:is_enabled;->;comment:'0=禁用,1=启用'"`
	IsDeleted int `gorm:"column:is_deleted;->;comment:'是否已删除'"`
}

func (c *ExtensionView) TableName() string {
	return "rapid.rx_market.extension_view"
}

func (c *ExtensionView) BeforeSave() error {
	return errors.New("视图不允许写操作")
}

// CreateView 创建 rapid.rx_market.extension_version_view 视图
func (c *ExtensionView) CreateView(db *gorm.DB) error {
	createViewSql := fmt.Sprintf(`
		CREATE VIEW %s AS (
			 SELECT extension.extension_id,
				extension.extension_uuid,
				extension.extension_name,
				extension.use_version,
				extension.script_hash,
				extension.metadata,
				extension.created_time AS extension_created_time,
				extension.updated_time AS extension_updated_time,
				extension.creator_id AS extension_creator_id,
				extension_create_user.username AS extension_creator_username,
				extension.updater_id AS extension_updater_id,
				extension_update_user.username AS extension_updater_username,
				extension.description AS extension_description,
				extension_group.extension_group_id,
				extension_group.extension_group_uuid,
				extension_group.extension_group_name,
				extension_version.extension_version_id,
				extension_version.script_content,
				extension_version.version,
				extension_version.description AS extension_version_description,
				extension_version.created_time AS extension_version_created_time,
				extension_version.updated_time AS extension_version_updated_time,
				extension_version.creator_id AS extension_version_creator_id,
				extension_create_user.username AS extension_version_creator_username,
				extension_version.updater_id AS extension_version_updater_id,
				extension_update_user.username AS extension_version_updater_username,
				CASE
					WHEN extension_group.enabled = 0 THEN 0
					WHEN extension.enabled = 0 THEN 0
					ELSE 1
				END AS is_enabled,
				CASE
					WHEN COALESCE((extension_group.status ->> 'is_deleted'::text)::boolean, false) = false THEN 0
					WHEN COALESCE((extension.status ->> 'is_deleted'::text)::boolean, false) = false THEN 0
					ELSE 1
				END AS is_deleted
		     FROM rx_market.extension
				 LEFT JOIN rx_market.extension_version ON extension.use_version = extension_version.extension_version_id
				 LEFT JOIN rx_market.extension_group ON extension.extension_group_id = extension_group.extension_group_id
				 LEFT JOIN client."user" extension_create_user ON extension.creator_id = extension_create_user.user_id
				 LEFT JOIN client."user" extension_update_user ON extension.updater_id = extension_update_user.user_id
				 LEFT JOIN client."user" extension_version_create_user ON extension.creator_id = extension_version_create_user.user_id
				 LEFT JOIN client."user" extension_version_update_user ON extension.updater_id = extension_version_update_user.user_id
			WHERE 	
				extension.use_version IS NOT NULL AND 
				extension.script_hash IS NOT NULL AND 
				extension_version.extension_version_id IS NOT NULL
		)		
	`, c.TableName())

	return db.Exec(createViewSql).Error
}
