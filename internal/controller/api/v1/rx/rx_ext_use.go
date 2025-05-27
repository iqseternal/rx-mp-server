package v1RX

import (
	"gorm.io/gorm/clause"
	"rx-mp/internal/biz"
	rdMarket "rx-mp/internal/models/rd/rx_market"
	"rx-mp/internal/pkg/rx"
	"rx-mp/internal/pkg/storage"
)

type UseExtensionQuery struct {
	ExtensionId   int64  `form:"extension_id" binding:"required,gt=0"`
	ExtensionUuid string `form:"extension_uuid" binding:"required,uuid"`
}

// UseExtension public: 对接某个扩展
func UseExtension(c *rx.Context) {
	var query UseExtensionQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	type Extension struct {
		ExtensionID        int64       `json:"extension_id"`
		ExtensionUuid      string      `json:"extension_uuid"`
		ExtensionName      string      `json:"extension_name"`
		Metadata           interface{} `json:"metadata"`
		UseVersion         *int64      `json:"use_version"`
		ScriptHash         *string     `json:"script_hash"`
		ExtensionGroupID   int64       `json:"extension_group_id"`
		ExtensionGroupUuid string      `json:"extension_group_uuid"`
		ExtensionGroupName string      `json:"extension_group_name"`
		ExtensionVersionID int64       `json:"extension_version_id"`
		ScriptContent      *string     `json:"script_content"`
		Version            int64       `json:"version"`
	}

	var extension Extension
	result := storage.RdPostgres.
		Model(&rdMarket.ExtensionView{}).
		Where("is_deleted = ?", 0).
		Where("is_enabled = ?", 1).
		Where("extension_id = ?", query.ExtensionId).
		Where("extension_uuid = ?", query.ExtensionUuid).
		Select(
			"extension_id",
			"extension_uuid",
			"extension_name",
			"metadata",
			"use_version",
			"script_hash",
			"extension_group_id",
			"extension_group_uuid",
			"extension_group_name",
			"extension_version_id",
			"script_content",
			"version",
		).
		First(&extension)

	if result.Error != nil {
		c.FailWithMessage("扩展不存在", nil)
		return
	}

	c.Ok(extension)
}

// UseExtensionGroupQuery 定义 UseExtensionGroup 接口的请求参数结构体
type UseExtensionGroupQuery struct {
	ExtensionGroupId   int64  `form:"extension_group_id" binding:"required,gt=0"`
	ExtensionGroupUuid string `form:"extension_group_uuid" binding:"required,uuid"`
}

// UseExtensionGroup public: 对接某个插件组
func UseExtensionGroup(c *rx.Context) {
	var query UseExtensionGroupQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	type Extension struct {
		ExtensionID        int64       `json:"extension_id"`
		ExtensionUuid      string      `json:"extension_uuid"`
		ExtensionName      string      `json:"extension_name"`
		Metadata           interface{} `json:"metadata"`
		UseVersion         *int64      `json:"use_version"`
		ScriptHash         *string     `json:"script_hash"`
		ExtensionGroupID   int64       `json:"extension_group_id"`
		ExtensionGroupUuid string      `json:"extension_group_uuid"`
		ExtensionGroupName string      `json:"extension_group_name"`
		ExtensionVersionID int64       `json:"extension_version_id"`
		ScriptContent      *string     `json:"script_content"`
		Version            int64       `json:"version"`
	}

	var extensionList []Extension
	result := storage.RdPostgres.
		Model(&rdMarket.ExtensionView{}).
		Where("is_deleted = ?", 0).
		Where("is_enabled = ?", 1).
		Where("extension_group_id = ?", query.ExtensionGroupId).
		Where("extension_group_uuid = ?", query.ExtensionGroupUuid).
		Select(
			"extension_id",
			"extension_uuid",
			"extension_name",
			"metadata",
			"use_version",
			"script_hash",
			"extension_group_id",
			"extension_group_uuid",
			"extension_group_name",
			"extension_version_id",
			"script_content",
			"version",
		).
		Find(&extensionList)

	if result.Error != nil {
		c.FailWithCode(biz.DatabaseQueryError, nil)
		return
	}

	c.Ok(extensionList)
}

type UseExtensionHeartbeatVoucher struct {
	ExtensionId   *int64  `json:"extension_id" binding:"required,gt=0"`
	ExtensionUuid *string `json:"extension_uuid" binding:"required,uuid"`
	ScriptHash    *string `json:"script_hash" binding:"required"`
}

// UseExtensionHeartbeatPayload 定义 UseExtensionHeartbeat 接口的请求参数结构体
type UseExtensionHeartbeatPayload struct {
	Vouchers []UseExtensionHeartbeatVoucher `json:"vouchers" binding:"required,min=1"`
}

// UseExtensionHeartbeat public: 插件心跳
func UseExtensionHeartbeat(c *rx.Context) {
	var payload UseExtensionHeartbeatPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var orConditions []clause.Expression
	for _, voucher := range payload.Vouchers {
		if voucher.ExtensionId == nil || voucher.ExtensionUuid == nil || voucher.ScriptHash == nil {
			c.FailWithCodeMessage(biz.ParameterError, "missing required fields", nil)
			return
		}

		orConditions = append(orConditions, clause.And(
			clause.Eq{Column: "extension_id", Value: *voucher.ExtensionId},
			clause.Eq{Column: "extension_uuid", Value: *voucher.ExtensionUuid},
			clause.Neq{Column: "script_hash", Value: *voucher.ScriptHash},
		))
	}

	var extensionIDs []int64
	result := storage.RdPostgres.
		Model(&rdMarket.Extension{}).
		Select("extension_id").
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false).
		Where("enabled = ?", 1).
		Where(clause.Or(orConditions...)).
		Pluck("extension_id", &extensionIDs)

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	c.Ok(extensionIDs)
}
