package config

import (
	"famcache/domain/config"
	"flag"
)

type Config struct {
	Port string
}

func (c *Config) GetPort() string {
	return c.Port
}

// NewConfig creates a new configuration
// Reads the port from command line arguments
func NewConfig() config.FamcacheConfig {
	port := flag.String("port", "3577", "Port to run the server on")
	flag.Parse()

	if *port == "" {
		*port = "3577"
	}

	return &Config{
		Port: *port,
	}
}
