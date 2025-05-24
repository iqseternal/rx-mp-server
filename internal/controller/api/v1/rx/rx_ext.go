package v1RX

import (
	"gorm.io/gorm"
	"rx-mp/internal/biz"
	rdMarket "rx-mp/internal/models/rd/rx_market"
	"rx-mp/internal/pkg/mbic"
	"rx-mp/internal/pkg/rx"
	"rx-mp/internal/pkg/storage"
)

type GetExtensionListQuery struct {
	ExtensionGroupId *int64 `form:"extension_group_id" binding:"omitempty,gt=0"`

	ExtensionId   *int64  `form:"extension_id" binding:"omitempty,gt=0"`
	ExtensionUuid *string `form:"extension_uuid" binding:"omitempty"`
	ExtensionName *string `form:"extension_name" binding:"omitempty"`

	Page     *int `form:"page" binding:"omitempty,gt=0"`
	PageSize *int `form:"page_size" binding:"omitempty,gt=0"`
}

func GetExtensionList(c *rx.Context) {
	var query GetExtensionListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.FailWithMessage(err.Error(), nil)
		return
	}

	var total int64
	var extensionList []rdMarket.Extension
	db := storage.RdPostgres.Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false)

	if query.ExtensionGroupId != nil {
		db = db.Where("extension_group_id = ?", *query.ExtensionGroupId)
	}

	if query.ExtensionId != nil {
		db = db.Where("extension_id = ?", *query.ExtensionId)
	}

	if query.ExtensionUuid != nil {
		db = db.Where("extension_uuid = ?", *query.ExtensionUuid)
	}

	if query.ExtensionName != nil {
		db = db.Where("extension_name like ?", "%"+*query.ExtensionName+"%")
	}

	db = db.Order("created_time desc")
	if err := db.Model(&rdMarket.Extension{}).Count(&total).Error; err != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, err.Error(), nil)
		return
	}

	if query.Page != nil && query.PageSize != nil {
		db = db.Offset((*query.Page - 1) * *query.PageSize).Limit(*query.PageSize)
	}

	result := db.Find(&extensionList)

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	c.Ok(&rx.H{
		"list":  extensionList,
		"total": total,
	})
}

type AddExtensionPayload struct {
	ExtensionGroupId   int64   `json:"extension_group_id"`
	ExtensionGroupUuid *string `json:"extension_group_uuid" binding:"omitempty"`

	ExtensionName string `json:"extension_name" binding:"required"`

	Description *string `json:"description" binding:"omitempty"`
}

// AddExtension 添加一个插件, 默认直接创建一个版本, 然后再创建一个活跃的插件
func AddExtension(c *rx.Context) {
	user, err := mbic.GetMBICUser(c.Context)
	if err != nil {
		c.FailWithCodeMessage(biz.MBICQueryError, err.Error(), nil)
		return
	}

	var payload AddExtensionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var extensionGroup rdMarket.ExtensionGroup
	db := storage.RdPostgres.Where("extension_group_id = ?", payload.ExtensionGroupId)

	if payload.ExtensionGroupUuid != nil {
		db = db.Where("extension_group_uuid = ?", *payload.ExtensionGroupUuid)
	}

	extensionGroupResult := db.First(&extensionGroup)

	if extensionGroupResult.Error != nil {
		c.FailWithMessage("扩展组无效", nil)
		return
	}

	result := storage.RdPostgres.Create(&rdMarket.Extension{
		ExtensionGroupId: payload.ExtensionGroupId,
		ExtensionName:    payload.ExtensionName,
		CreatorID:        &user.UserID,
		Description:      payload.Description,
	})

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	c.Ok(nil)
}

type DelExtensionPayload struct {
	ExtensionId   int64  `json:"extension_id" binding:"required"`
	ExtensionUuid string `json:"extension_uuid" binding:"required"`
}

func DelExtension(c *rx.Context) {
	user, err := mbic.GetMBICUser(c.Context)
	if err != nil {
		c.FailWithCodeMessage(biz.MBICQueryError, err.Error(), nil)
		return
	}

	var payload DelExtensionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	updates := map[string]interface{}{
		"updater_id": user.UserID,
		"status": gorm.Expr(
			"jsonb_set(status, '{is_deleted}', ?)",
			true,
		),
	}

	result := storage.RdPostgres.
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false).
		Where("extension_id = ?", payload.ExtensionId).
		Where("extension_uuid = ?", payload.ExtensionUuid).
		Updates(updates)

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
	ExtensionId   int64  `form:"extension_id" binding:"required"`
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
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false).
		Where("extension_id = ?", query.ExtensionId).
		Where("extension_uuid = ?", query.ExtensionUuid).
		First(&extension)

	if result.Error != nil {
		c.FailWithMessage("扩展不存在", nil)
		return
	}

	c.Ok(extension)
}

type ModifyExtensionPayload struct {
	ExtensionId   int64  `json:"extension_id" binding:"required"`
	ExtensionUuid string `json:"extension_uuid" binding:"required"`

	ExtensionName *string      `json:"extension_name" binding:"omitempty"`
	Metadata      *interface{} `json:"metadata" binding:"omitempty"`
	Enabled       *int         `json:"enabled" binding:"omitempty"`
	UseVersion    *int         `json:"use_version" binding:"omitempty"`
}

func ModifyExtension(c *rx.Context) {
	user, err := mbic.GetMBICUser(c.Context)
	if err != nil {
		c.FailWithCodeMessage(biz.MBICQueryError, err.Error(), nil)
		return
	}

	var payload ModifyExtensionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	// 动态构建更新字段
	updates := make(map[string]interface{})
	updates["updater_id"] = &user.UserID

	if payload.ExtensionName != nil {
		updates["extension_name"] = *payload.ExtensionName
	}

	if payload.Metadata != nil {
		updates["metadata"] = *payload.Metadata
	}

	if payload.Enabled != nil {
		updates["enabled"] = payload.Enabled
	}

	if payload.UseVersion != nil {
		updates["use_version"] = *payload.UseVersion
	}

	if len(updates) == 0 {
		c.FailWithCode(biz.AttemptUpdateInValidData, nil)
		return
	}

	result := storage.RdPostgres.
		Model(&rdMarket.Extension{}).
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false).
		Where("extension_id = ?", payload.ExtensionId).
		Where("extension_uuid = ?", payload.ExtensionUuid).
		Updates(updates) // 一次性更新所有字段

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	if result.RowsAffected == 0 {
		c.FailWithCode(biz.AttemptUpdateInValidData, nil)
		return
	}

	c.Ok(nil)
}
