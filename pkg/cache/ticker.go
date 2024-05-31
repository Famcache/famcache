package cache

import "time"

func (c *Cache) CleanupExpired() {
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

			// Check if the stored value has expired
			if elapsed >= int64(*storedValue.TTL)/1000 {
				delete(c.storage, key)
			}
		}
	}
}
