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
	ExtensionId int64 `json:"extension_id" binding:"required"`

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

	var versions int64
	if err := storage.RdPostgres.Model(&rdMarket.ExtensionVersion{}).Where("extension_id = ?", payload.ExtensionId).Count(&versions).Error; err != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, err.Error(), nil)
		return
	}

	result := storage.RdPostgres.Create(&rdMarket.ExtensionVersion{
		ExtensionId:   payload.ExtensionId,
		Version:       versions + 1,
		ScriptContent: payload.ScriptContent,
		Description:   *payload.Description,
		CreatorID:     &user.UserID,
	})

	if result.Error != nil {
		c.FailWithCodeMessage(biz.DatabaseQueryError, result.Error.Error(), nil)
		return
	}

	c.Ok(result.Row())
}
