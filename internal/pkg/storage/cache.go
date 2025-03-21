package storage

import (
	"rx-mp/pkg/storage"
)

var MemoCache *storage.MemoryCache

func init() {
	cacheSize := 100 * 1024 * 1024 // 100MB

	MemoCache = storage.NewMemoryCache(cacheSize)
}
