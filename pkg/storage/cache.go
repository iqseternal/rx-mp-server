package storage

import (
	"encoding/json"

	"github.com/coocood/freecache"
)

// MemoryCache 封装 FreeCache，支持字符串操作
type MemoryCache struct {
	cache *freecache.Cache
}

// NewMemoryCache 创建一个新的 new MemoryCache
func NewMemoryCache(size int) *MemoryCache {
	return &MemoryCache{
		cache: freecache.NewCache(size),
	}
}

// Set 设置缓存值
func (sc *MemoryCache) Set(key, value string, expireSeconds int) error {
	return sc.cache.Set([]byte(key), []byte(value), expireSeconds)
}

// Get 获取缓存值
func (sc *MemoryCache) Get(key string) (string, error) {
	value, err := sc.cache.Get([]byte(key))

	if err != nil {
		return "", err
	}
	return string(value), nil
}

// GetWithStruct 获取一个结构体
func (sc *MemoryCache) GetWithStruct(key string, v interface{}) error {
	value, err := sc.cache.Get([]byte(key))

	if err != nil {
		return err
	}

	err = json.Unmarshal(value, v)

	if err != nil {
		return err
	}

	return nil
}

// Delete 删除缓存值
func (sc *MemoryCache) Delete(key string) bool {
	return sc.cache.Del([]byte(key))
}

// Clear 清空缓存
func (sc *MemoryCache) Clear() {
	sc.cache.Clear()
}
