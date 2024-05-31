package server

import (
	"famcache/domain/cache"
	"famcache/domain/logger"
	"famcache/domain/server"
	"famcache/pkg/actor"
	"net"
)

// Server is a struct that represents a server
type Server struct {
	listener net.Listener
	port     string
	logger   logger.Logger
	cache    cache.Cache
	actor    actor.Actor
}

// NewServer creates a new server
func NewServer(options ServerOptions) server.FamcacheServer {
	actors := actor.NewActor(&options.Logger, &options.Cache)

	return &Server{
		port:   options.Port,
		logger: options.Logger,
		cache:  options.Cache,
		actor:  actors,
	}
}
