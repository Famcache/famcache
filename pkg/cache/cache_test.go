package cache

import (
	"famcache/domain/cache"
	"testing"
)

func TestNewCache(t *testing.T) {
	c := NewCache()

	if c == nil {
		t.Error("Expected cache instance, got nil")
	}

	if _, ok := c.(*Cache); !ok {
		t.Error("Expected Cache instance, got different type")
	}

	t.Log("NewCache() test passed")
}

func TestCachePersistance(t *testing.T) {
	c := NewCache()

	const persistentKey = "persistentKey"
	const persistentValue = "persistantValue"

	if val, err := c.Get(persistentKey); err != nil && val != "" {
		t.Errorf("Expected error, got %s", val)
	}

	err := c.Set(cache.SetOptions{
		Key:   persistentKey,
		Value: persistentValue,
		TTL:   nil,
	})

	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if val, err := c.Get(persistentKey); err != nil && val != persistentValue {
		t.Errorf("Expected value, got %s", val)
	}

	err = c.Delete(persistentKey)

	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if val, err := c.Get(persistentKey); err != nil && val != "" {
		t.Errorf("Expected error, got %s", val)
	}
}
