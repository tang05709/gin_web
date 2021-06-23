package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	instance *cache.Cache
	once     sync.Once
)

//获取缓存单例
func Instance() *cache.Cache {
	once.Do(func() {
		c := cache.New(5*time.Minute, 10*time.Minute)
		instance = c
	})

	return instance
}
