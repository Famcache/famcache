package config

import "flag"

type Config struct {
	Port string
}

// NewConfig creates a new configuration
// Reads the port from command line arguments
func NewConfig() *Config {
	port := flag.String("port", "3577", "Port to run the server on")
	flag.Parse()

	if *port == "" {
		*port = "3577"
	}

	return &Config{
		Port: *port,
	}
}
