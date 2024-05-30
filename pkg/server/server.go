package server

import (
	"famcache/domain/cache"
	"famcache/domain/logger"
	"famcache/domain/server"
	"famcache/pkg/server/peers"
	"net"
)

// Server is a struct that represents a server
type Server struct {
	listener net.Listener
	port     string
	logger   logger.Logger
	cache    cache.Cache
	peers    []peers.Peer
}

// NewServer creates a new server
func NewServer(options ServerOptions) server.FamcacheServer {
	return &Server{
		port:   options.Port,
		logger: options.Logger,
		cache:  options.Cache,
		peers:  make([]peers.Peer, 0),
	}
}
