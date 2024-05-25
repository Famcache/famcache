package cache

import "famcache/domain/cache"

type Cache struct {
	storage map[string]string
}

func NewCache() cache.Cache {
	return &Cache{
		storage: make(map[string]string),
	}
}
