package v1RX

import (
	"rx-mp/internal/biz"
	"rx-mp/internal/pkg/rx"
)

// UseExtension public: 对接某个扩展
func UseExtension(c *rx.Context) {

}

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
}
