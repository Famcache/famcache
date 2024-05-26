package cache

import (
	"famcache/domain/cache"
)

func (c *Cache) Start() error {
	go c.CleanupExpired()

	return nil
}

func (c *Cache) Get(key string) (string, error) {
	val, ok := c.storage[key]

	if !ok {
		return "", nil
	}

	return val.Value, nil
}

func (c *Cache) Set(options cache.SetOptions) error {
	stored, err := NewStoredValue(options.Value, options.TTL)

	if err != nil {
		return err
	}

	c.storage[options.Key] = stored

	return nil
}

func (c *Cache) Delete(key string) error {
	delete(c.storage, key)

	return nil
}
