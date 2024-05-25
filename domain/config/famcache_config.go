package config

// FamcacheConfig is an interface that defines the configuration for the Famcache server
type FamcacheConfig interface {
	GetPort() string
}
