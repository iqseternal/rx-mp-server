package v1RX

import (
	"rx-mp/internal/biz"
	rdMarket "rx-mp/internal/models/rd/rx_market"
	"rx-mp/internal/pkg/mbic"
	"rx-mp/internal/pkg/rx"
	"rx-mp/internal/pkg/storage"
)

type GetGetExtensionVersionListQuery struct {
	ExtensionId   int64  `form:"extension_id" binding:"required"`
	ExtensionUuid string `form:"extension_uuid" binding:"required"`

	Page     *int `form:"page" binding:"omitempty,gt=0"`
	PageSize *int `form:"page_size" binding:"omitempty,gt=0"`

	ExtensionVersionId *int64  `form:"extension_version_id" binding:"omitempty"`
	Description        *string `form:"description" binding:"omitempty,max=255"`
	UpdaterID          *int64  `form:"updater_id" binding:"omitempty"`
}

func GetExtensionVersionList(c *rx.Context) {
	var query GetGetExtensionVersionListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.FailWithCode(biz.ParameterError, nil)
		return
	}

	var total int64
	var extensionVersionList []rdMarket.ExtensionVersion

	db := storage.RdPostgres.Model(&rdMarket.ExtensionVersion{})

	db = db.Where("extension_id = ?", query.ExtensionId)

	if query.ExtensionVersionId != nil {
		db = db.Where("extension_version_id = ?", query.ExtensionVersionId)
	}

	if query.Description != nil {
		db = db.Where("description like '%?%'", query.Description)
	}

	if query.UpdaterID != nil {
		db = db.Where("updater_id = ?", query.UpdaterID)
	}

	if err := db.Model(&rdMarket.ExtensionVersion{}).Count(&total).Error; err != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, err.Error(), nil)
		return
	}

	db = db.Order("created_time desc")

	if query.Page != nil && query.PageSize != nil {
		db = db.Offset((*query.Page - 1) * *query.PageSize).Limit(*query.PageSize)
	}

	if err := db.Find(&extensionVersionList).Error; err != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, err.Error(), nil)
		return
	}

	c.Ok(&rx.H{
		"list":  extensionVersionList,
		"total": total,
	})
}

type AddExtensionVersionPayload struct {
	ExtensionId   int64  `json:"extension_id" binding:"required"`
	ExtensionUuid string `json:"extension_uuid" binding:"required"`

	ScriptContent string  `json:"script_content" binding:"required"`
	Description   *string `json:"description" binding:"omitempty,max=255"`
}

func AddExtensionVersion(c *rx.Context) {
	user, err := mbic.GetMBICUser(c.Context)
	if err != nil {
		c.FailWithCodeMessage(biz.MBICQueryError, err.Error(), nil)
		return
	}

	var payload AddExtensionVersionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var extension rdMarket.Extension
	extResult := storage.RdPostgres.Model(&rdMarket.Extension{}).
		Where("extension_id = ?", payload.ExtensionId).
		Where("extension_uuid = ?", payload.ExtensionUuid).
		First(&extension)

	if extResult.Error != nil {
		c.FailWithMessage("无效扩展", nil)
		return
	}

	var versions int64
	extVersionCountResult := storage.RdPostgres.Model(&rdMarket.ExtensionVersion{}).
		Where("extension_id = ?", payload.ExtensionId).
		Count(&versions)

	if extVersionCountResult.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, extVersionCountResult.Error.Error(), nil)
		return
	}

	var newExtensionVersion = rdMarket.ExtensionVersion{
		ExtensionId:   payload.ExtensionId,
		Version:       versions + 1,
		ScriptContent: payload.ScriptContent,
		CreatorID:     &user.UserID,
	}

	if payload.Description != nil {
		newExtensionVersion.Description = *payload.Description
	}

	result := storage.RdPostgres.Create(&newExtensionVersion)

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	c.Ok(result.Row())
}

type ApplyExtensionVersionPayload struct {
	ExtensionId        int64  `json:"extension_id" binding:"required"`
	ExtensionUuid      string `json:"extension_uuid" binding:"required"`
	ExtensionVersionId int64  `json:"extension_version_id" binding:"required"`
}

func ApplyExtensionVersion(c *rx.Context) {
	user, err := mbic.GetMBICUser(c.Context)
	if err != nil {
		c.FailWithCodeMessage(biz.MBICQueryError, err.Error(), nil)
		return
	}

	var payload ApplyExtensionVersionPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.FailWithCodeMessage(biz.ParameterError, err.Error(), nil)
		return
	}

	var extension rdMarket.Extension
	extResult := storage.RdPostgres.Model(&rdMarket.Extension{}).
		Where("extension_id = ?", payload.ExtensionId).
		Where("extension_uuid = ?", payload.ExtensionUuid).
		First(extension)

	if extResult.Error != nil {
		c.FailWithMessage("扩展不存在", nil)
		return
	}

	var extensionVersion rdMarket.ExtensionVersion
	extVersionResult := storage.RdPostgres.Model(&rdMarket.ExtensionVersion{}).
		Where("extension_version_id = ?", payload.ExtensionVersionId).
		First(extensionVersion)

	if extVersionResult.Error != nil {
		c.FailWithMessage("扩展版本不存在", nil)
		return
	}

	updates := map[string]interface{}{
		"use_version": extensionVersion.Version,
		"updater_id":  user.UserID,
	}

	useVersionResult := storage.RdPostgres.Model(&rdMarket.Extension{}).
		Where("extension_id = ?", payload.ExtensionId).
		Where("extension_uuid = ?", payload.ExtensionUuid).
		Updates(updates)

	if useVersionResult.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, useVersionResult.Error.Error(), nil)
		return
	}

	c.Ok(nil)
}
