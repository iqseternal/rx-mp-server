package v1RX

import (
	"gorm.io/gorm"
	"net/http"
	"rx-mp/internal/biz"
	rdMarket "rx-mp/internal/models/rd/rx_market"
	"rx-mp/internal/pkg/rx"
	"rx-mp/internal/pkg/storage"
)

type AddExtensionPayload struct {
	ExtensionGroupId   int    `json:"extension_group_id"`
	ExtensionGroupUuid string `json:"extension_group_uuid"`

	ExtensionName string `json:"extension_name"`
}

// AddExtension 添加一个插件, 默认直接创建一个版本, 然后再创建一个活跃的插件
func AddExtension(c *rx.Context) {
	var payload AddExtensionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var extensionGroup rdMarket.ExtensionGroup
	extensionGroupResult := storage.RdPostgres.
		Where("extension_group_id = ?", payload.ExtensionGroupId).
		Where("extension_group_uuid = ?", payload.ExtensionGroupUuid).
		First(&extensionGroup)

	if extensionGroupResult.Error != nil {
		c.FailWithMessage("扩展组无效", nil)
		return
	}

	result := storage.RdPostgres.Create(&rdMarket.Extension{
		ExtensionGroupId: payload.ExtensionGroupId,
		ExtensionName:    payload.ExtensionName,
	})

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	c.Ok(nil)
}

type DelExtensionPayload struct {
	ExtensionId   int    `json:"extension_id" binding:"required"`
	ExtensionUuid string `json:"extension_uuid" binding:"required"`
}

func DelExtension(c *rx.Context) {
	var payload DelExtensionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	result := storage.RdPostgres.
		Where("extension_id = ?", payload.ExtensionId).
		Where("extension_uuid = ?", payload.ExtensionUuid).
		Update("status", gorm.Expr(
			"jsonb_set(status, '{is_deleted}', ?)",
			true,
		))

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	if result.RowsAffected == 0 {
		c.FailWithCode(biz.AttemptDeleteInValidData, nil)
		return
	}

	c.Ok(nil)
}

type GetExtensionQuery struct {
	ExtensionId   int    `form:"extension_id" binding:"required"`
	ExtensionUuid string `form:"extension_uuid" binding:"required"`
}

func GetExtension(c *rx.Context) {
	var query GetExtensionQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var extension rdMarket.Extension
	result := storage.RdPostgres.
		Where("extension_id = ?", query.ExtensionId).
		Where("extension_uuid = ?", query.ExtensionUuid).
		First(&extension)

	if result.Error != nil {
		c.FailWithMessage("扩展不存在", nil)
		return
	}

	c.Ok(extension)
}

func ModifyExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}

func ActiveExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}

func DeactiveExtension(c *rx.Context) {
	c.Finish(http.StatusMethodNotAllowed, &rx.R{
		Code:    biz.NotImplemented,
		Message: biz.Message(biz.NotImplemented),
		Data:    nil,
	})
}
