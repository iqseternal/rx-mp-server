package rdMarket

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// ExtensionVersionView 插件版本视图
type ExtensionVersionView struct {
	// 基础版本字段
	ExtensionVersionId int64  `gorm:"primaryKey;column:extension_version_id;autoIncrement:false"`
	ScriptContent      string `gorm:"column:script_content;type:text;serializer:compress"`
	Version            int64  `gorm:"column:version;index:idx_version;comment:'版本号排序使用'"`
	Description        string `gorm:"column:description;type:varchar(500)"`

	// 关联扩展插件
	ExtensionID   int64       `gorm:"column:extension_id;index:idx_extension"`
	ExtensionUUID string      `gorm:"column:extension_uuid;type:uuid;not null"`
	ExtensionName string      `gorm:"column:extension_name;index;type:varchar(255)"`
	Metadata      interface{} `gorm:"column:metadata;type:jsonb;serializer:json"`

	// 关联扩展组
	ExtensionGroupID   int64  `gorm:"column:extension_group_id"`
	ExtensionGroupUUID string `gorm:"column:extension_group_uuid;type:uuid"`
	ExtensionGroupName string `gorm:"column:extension_group_name;type:varchar(255)"`

	// 状态计算字段（只读）
	EnabledStatus int `gorm:"column:enabled_status;->;comment:'0=禁用,1=启用'"`
	StatusResult  int `gorm:"column:status_result;->;comment:'0=删除,1=正常'"`

	// 用户关联（指针处理空值）
	CreatorID       *uint64 `gorm:"column:creator_id"`
	CreatorUsername string  `gorm:"->;column:creator_username;type:varchar(64)"`
	UpdaterID       *uint64 `gorm:"column:updater_id"`
	UpdaterUsername string  `gorm:"->;column:updater_username;type:varchar(64)"`

	// 自动时间戳（与基表同步）
	CreatedTime time.Time `gorm:"column:created_time;->"`
	UpdatedTime time.Time `gorm:"column:updated_time;->"`
}

func (c *ExtensionVersionView) TableName() string {
	return "rapid.rx_market.extension_version_view"
}

func (c *ExtensionVersionView) BeforeSave(db *gorm.DB) error {
	return errors.New("视图不允许写操作")
}

// CreateView 创建 rapid.rx_market.extension_version_view 视图
func (c *ExtensionVersionView) CreateView(db *gorm.DB) error {
	createViewSql := fmt.Sprintf(`
		CREATE VIEW %s AS (
			SELECT 
				extension_version.extension_version_id,
				extension_version.script_content,
				extension_version.version,
				extension_version.description,
				extension_version.created_time,
				extension_version.updated_time,
				extension_version.creator_id,
				create_user.username AS creator_username,
				extension_version.updater_id,
				update_user.username AS updater_username,
				extension.extension_id,
				extension.extension_uuid,
				extension.extension_name,
				extension.metadata,
				extension_group.extension_group_id,
				extension_group.extension_group_uuid,
				extension_group.extension_group_name,
				CASE
					WHEN extension_group.enabled = 0 THEN 0
					WHEN extension.enabled = 0 THEN 0
					ELSE 1
				END AS enabled_status,
				CASE
					WHEN extension_group.status ? 'is_deleted'::text AND (extension_group.status ->> 'is_deleted'::text) = 'true'::text THEN 0
					WHEN extension.status ? 'is_deleted'::text AND (extension.status ->> 'is_deleted'::text) = 'true'::text THEN 0
					ELSE 1
				END AS status_result
			FROM rx_market.extension_version
				LEFT JOIN rx_market.extension ON extension_version.extension_id = extension.extension_id
				LEFT JOIN rx_market.extension_group ON extension.extension_group_id = extension_group.extension_group_id
				LEFT JOIN client."user" create_user ON extension_version.creator_id = create_user.user_id
				LEFT JOIN client."user" update_user ON extension_version.updater_id = update_user.user_id
		)		
	`, c.TableName())

	return db.Exec(createViewSql).Error
}
