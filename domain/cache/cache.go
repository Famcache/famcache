package cache

// Cache is an interface that defines the methods for a cache
type Cache interface {
	Get(key string) (string, error)
	Set(options SetOptions) error
	Delete(key string) error
}
