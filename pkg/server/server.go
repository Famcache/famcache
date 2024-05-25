package server

import (
	"famcache/domain/cache"
	"famcache/domain/logger"
	"famcache/domain/server"
	"net"
)

// Server is a struct that represents a server
type Server struct {
	listener net.Listener
	port     string
	logger   logger.Logger
	cache    cache.Cache
}

// NewServer creates a new server
func NewServer(options ServerOptions) server.FamcacheServer {
	return &Server{
		port:   options.Port,
		logger: options.Logger,
		cache:  options.Cache,
	}
}
