package v1RX

import (
	"rx-mp/internal/biz"
	rdMarket "rx-mp/internal/models/rd/rx_market"
	"rx-mp/internal/pkg/rx"
	"rx-mp/internal/pkg/storage"
)

type UseExtensionPayload struct {
	ExtensionId   int64  `json:"extension_id" binding:"required,gt=0"`
	ExtensionUuid string `json:"extension_uuid" binding:"required,uuid"`
}

// UseExtension public: 对接某个扩展
func UseExtension(c *rx.Context) {
	var payload UseExtensionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var extension rdMarket.Extension
	result := storage.RdPostgres.
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false).
		Where("extension_id = ?", payload.ExtensionId).
		Where("extension_uuid = ?", payload.ExtensionUuid).
		First(&extension)

	if result.Error != nil {
		c.FailWithMessage("扩展不存在", nil)
		return
	}

	c.Ok(extension)
}

// UseExtensionGroupPayload 定义 UseExtensionGroup 接口的请求参数结构体
type UseExtensionGroupPayload struct {
	ExtensionGroupId   int64  `json:"extension_group_id" binding:"required,gt=0"`
	ExtensionGroupUuid string `json:"extension_group_uuid" binding:"required,uuid"`
}

// UseExtensionGroup public: 对接某个插件组
func UseExtensionGroup(c *rx.Context) {
	var payload UseExtensionGroupPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var extensionGroup rdMarket.ExtensionGroup
	result := storage.RdPostgres.
		Model(&rdMarket.ExtensionGroup{}).
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false).
		Where("extension_group_id = ?", payload.ExtensionGroupId).
		Where("extension_group_uuid = ?", payload.ExtensionGroupUuid).
		First(&extensionGroup)

	if result.Error != nil {
		c.FailWithMessage("扩展组不存在", nil)
		return
	}

	var extensionList []rdMarket.Extension
	result = storage.RdPostgres.
		Model(&rdMarket.Extension{}).
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false).
		Where("extension_group_id = ?", payload.ExtensionGroupId).
		Find(&extensionList)

	if result.Error != nil {
		c.FailWithMessage("扩展列表不存在", nil)
		return
	}

	c.Ok(extensionList)
}
