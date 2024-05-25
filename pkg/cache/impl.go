package cache

import "famcache/domain/cache"

func (c *Cache) Get(key string) (string, error) {
	return c.storage[key], nil
}

func (c *Cache) Set(options cache.SetOptions) error {
	c.storage[options.Key] = options.Value

	return nil
}

func (c *Cache) Delete(key string) error {
	delete(c.storage, key)

	return nil
}
