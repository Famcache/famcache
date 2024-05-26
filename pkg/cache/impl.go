package cache

import (
	"famcache/domain/cache"
	"time"
)

func (c *Cache) Start() error {
	go func() {
		for range c.ticker.C {
			// Iterate over all keys in the storage
			for key := range c.storage {
				// Get the stored value
				storedValue := c.storage[key]

				if storedValue.TTL == nil {
					continue
				}

				// Calculate the time passed since the stored value was created
				elapsed := time.Now().Unix() - storedValue.CreatedAt

				println("Elapsed", elapsed, "TTL", storedValue.TTL)
				// Check if the stored value has expired
				if elapsed >= *storedValue.TTL/1000 {
					// Delete the stored value
					println("Deleting key", key)
					delete(c.storage, key)
				}
			}

		}
	}()

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
