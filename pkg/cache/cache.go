package cache

import (
	"famcache/domain/cache"
	"time"
)

type Cache struct {
	storage map[string]*StoredValue
	ticker  *time.Ticker
}

func NewCache() cache.Cache {
	return &Cache{
		storage: make(map[string]*StoredValue),
		ticker:  time.NewTicker(time.Millisecond * 100),
	}
}
