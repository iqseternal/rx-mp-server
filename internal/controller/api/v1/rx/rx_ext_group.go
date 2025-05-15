package v1RX

import (
	"gorm.io/gorm"
	"rx-mp/internal/biz"
	rdMarket "rx-mp/internal/models/rd/rx_market"
	"rx-mp/internal/pkg/mbic"
	"rx-mp/internal/pkg/rx"
	"rx-mp/internal/pkg/storage"
)

type GetGetExtensionGroupListQuery struct {
	ExtensionGroupId   *int64  `form:"extension_group_id" binding:"omitempty,gt=0"`
	ExtensionGroupName *string `form:"extension_group_name" binding:"omitempty"`

	Page     *int `form:"page" binding:"omitempty,gt=0"`
	PageSize *int `form:"page_size" binding:"omitempty,gt=0"`
}

func GetExtensionGroupList(c *rx.Context) {
	var query GetGetExtensionGroupListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.FailWithMessage(err.Error(), nil)
		return
	}

	var total int64
	var extensionGroupList []rdMarket.ExtensionGroup
	db := storage.RdPostgres.
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false)

	if query.ExtensionGroupId != nil {
		db = db.Where("extension_group_id = ?", *query.ExtensionGroupId)
	}

	if query.ExtensionGroupName != nil {
		db = db.Where("extension_group_name like ?", "%"+*query.ExtensionGroupName+"%")
	}

	if err := db.Model(&rdMarket.ExtensionGroup{}).Count(&total).Error; err != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, err.Error(), nil)
		return
	}

	db = db.Order("created_time desc")

	if query.Page != nil && query.PageSize != nil {
		db = db.Offset((*query.Page - 1) * *query.PageSize).Limit(*query.PageSize)
	}

	if err := db.Find(&extensionGroupList).Error; err != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, err.Error(), nil)
		return
	}

	c.Ok(&rx.H{
		"list":  extensionGroupList,
		"total": total,
	})
}

type AddExtensionGroupPayload struct {
	ExtensionGroupName string  `json:"extension_group_name" binding:"required"`
	Description        *string `json:"description" binding:"omitempty"`
}

func AddExtensionGroup(c *rx.Context) {
	user, err := mbic.GetMBICUser(c.Context)
	if err != nil {
		c.FailWithCodeMessage(biz.MBICQueryError, err.Error(), nil)
		return
	}

	var payload AddExtensionGroupPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var newExtensionGroup = &rdMarket.ExtensionGroup{
		ExtensionGroupName: payload.ExtensionGroupName,
		CreatorID:          &user.UserID,
	}

	if payload.Description != nil {
		newExtensionGroup.Description = *payload.Description
	}

	result := storage.RdPostgres.Create(newExtensionGroup)

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	c.Ok(result.Row())
}

type DelExtensionGroupCertificate struct {
	ExtensionGroupId   int64  `json:"extension_group_id" binding:"required,gt=0"`
	ExtensionGroupUuid string `json:"extension_group_uuid" binding:"required,uuid"`
}

type DelExtensionGroupPayload struct {
	Certificates []DelExtensionGroupCertificate `json:"certificates" binding:"required,gt=0,dive"`
}

func DelExtensionGroup(c *rx.Context) {
	user, err := mbic.GetMBICUser(c.Context)
	if err != nil {
		c.FailWithCodeMessage(biz.MBICQueryError, err.Error(), nil)
		return
	}

	var payload DelExtensionGroupPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	db := storage.RdPostgres
	var orConditions [][]interface{}
	for _, item := range payload.Certificates {
		orConditions = append(orConditions, []interface{}{item.ExtensionGroupId, item.ExtensionGroupUuid})
	}

	db = db.Model(&rdMarket.ExtensionGroup{}).
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false).
		Where("(extension_group_id, extension_group_uuid) in ?", orConditions)

	// 执行批量更新
	updates := map[string]interface{}{
		"updater_id": user.UserID,
		"status": gorm.Expr(
			"jsonb_set(status, '{is_deleted}', ?)",
			true,
		),
	}

	result := db.Updates(updates)

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	if result.RowsAffected == 0 {
		c.FailWithCodeMessage(biz.AttemptDeleteInValidData, "尝试删除无效数据", nil)
		return
	}

	c.Ok(nil)
}

type GetExtensionGroupQuery struct {
	ExtensionGroupId   *int64  `form:"extension_group_id" binding:"required"`
	ExtensionGroupUuid *string `form:"extension_group_uuid" binding:"required"`
}

func GetExtensionGroup(c *rx.Context) {
	var query GetExtensionGroupQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var extensionGroup rdMarket.ExtensionGroup
	result := storage.RdPostgres.
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false).
		Where("extension_group_id = ?", query.ExtensionGroupId).
		Where("extension_group_uuid = ?", query.ExtensionGroupUuid).
		First(&extensionGroup)

	if result.Error != nil {
		c.FailWithMessage("扩展组不存在", nil)
		return
	}

	c.Ok(extensionGroup)
}

type ModifyExtensionGroupPayload struct {
	ExtensionGroupId   int64  `json:"extension_group_id" binding:"required"`
	ExtensionGroupUuid string `json:"extension_group_uuid" binding:"required"`

	ExtensionGroupName *string `json:"extension_group_name" binding:"omitempty"`
	Description        *string `json:"description" binding:"omitempty"`
	Enabled            *int    `json:"enabled" binding:"omitempty"`
}

func ModifyExtensionGroup(c *rx.Context) {
	user, err := mbic.GetMBICUser(c.Context)
	if err != nil {
		c.FailWithCodeMessage(biz.MBICQueryError, err.Error(), nil)
		return
	}

	var payload ModifyExtensionGroupPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	// 动态构建更新字段
	updates := make(map[string]interface{})
	updates["updater_id"] = &user.UserID

	if payload.ExtensionGroupName != nil {
		updates["extension_group_name"] = *payload.ExtensionGroupName
	}

	if payload.Description != nil {
		updates["description"] = *payload.Description
	}

	if payload.Enabled != nil {
		updates["enabled"] = payload.Enabled
	}

	if len(updates) == 0 {
		c.FailWithCode(biz.AttemptUpdateInValidData, nil)
		return
	}

	result := storage.RdPostgres.
		Model(&rdMarket.ExtensionGroup{}).
		Where("COALESCE((status->>'is_deleted')::boolean, false) = ?", false).
		Where("extension_group_id = ?", payload.ExtensionGroupId).
		Where("extension_group_uuid = ?", payload.ExtensionGroupUuid).
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
