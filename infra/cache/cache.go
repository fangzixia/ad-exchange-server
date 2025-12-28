package cache

import (
	"errors"
	"sync"
	"time"
)

// ErrKeyNotFound 键不存在错误
var ErrKeyNotFound = errors.New("key not found")

// CacheItem 缓存项
type CacheItem struct {
	Value      interface{}
	ExpireTime time.Time
}

// MemoryCache 内存缓存
type MemoryCache interface {
	Get(key string) (interface{}, error)
	GetInt(key string) (int, error)
	Set(key string, value interface{}, expire time.Duration) error
	Delete(key string) error
}

// memoryCache 内存缓存实现
type memoryCache struct {
	items map[string]*CacheItem
	mu    sync.RWMutex
}

// NewMemoryCache 创建内存缓存实例
func NewMemoryCache() MemoryCache {
	c := &memoryCache{
		items: make(map[string]*CacheItem),
	}
	// 启动过期清理协程
	go c.cleanExpired()
	return c
}

// Get 获取缓存
func (c *memoryCache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	// 检查是否过期
	if time.Now().After(item.ExpireTime) {
		delete(c.items, key)
		return nil, ErrKeyNotFound
	}

	return item.Value, nil
}

// GetInt 获取int类型缓存
func (c *memoryCache) GetInt(key string) (int, error) {
	val, err := c.Get(key)
	if err != nil {
		return 0, err
	}

	intVal, ok := val.(int)
	if !ok {
		return 0, errors.New("value is not int type")
	}

	return intVal, nil
}

// Set 设置缓存
func (c *memoryCache) Set(key string, value interface{}, expire time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &CacheItem{
		Value:      value,
		ExpireTime: time.Now().Add(expire),
	}

	return nil
}

// Delete 删除缓存
func (c *memoryCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
	return nil
}

// cleanExpired 清理过期缓存
func (c *memoryCache) cleanExpired() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, item := range c.items {
			if time.Now().After(item.ExpireTime) {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
