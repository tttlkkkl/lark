package cache

import "time"

// Cache 简单缓存实现
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Del(key string) error
}
