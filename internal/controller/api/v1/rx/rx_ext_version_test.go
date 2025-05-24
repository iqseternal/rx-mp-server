package v1RX

import (
	rdMarket "rx-mp/internal/models/rd/rx_market"
	"rx-mp/internal/pkg/storage"
	"testing"
)

// TestSelectExtension tests the SelectExtension function
func TestSelectExtension(t *testing.T) {
	var extension *rdMarket.Extension
	extResult := storage.RdPostgres.Model(&rdMarket.Extension{}).
		Where("extension_id = ?", 42).
		Where("extension_uuid = ?", "a764c09c-7955-425f-8260-f44bcef4c978").
		First(&extension)

	if extResult.Error != nil {
		t.Error(extResult.Error.Error())
		return
	}

	r2 := storage.RdPostgres.Model(&rdMarket.Extension{}).Exec("SELECT * FROM \"rapid\".\"rx_market\".\"extension\" WHERE extension_id = 42 AND extension_uuid = 'a764c09c-7955-425f-8260-f44bcef4c978' ORDER BY \"rapid\".\"rx_market\".\"extension\".\"extension_id\" LIMIT 1")
	if r2.Error != nil {
		t.Error(r2.Error.Error())
	}
}
