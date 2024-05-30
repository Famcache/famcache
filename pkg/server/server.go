package server

import (
	"famcache/domain/cache"
	"famcache/domain/logger"
	"famcache/domain/pubsub"
	"famcache/domain/server"
	"famcache/pkg/server/peers"
	queue "famcache/pkg/server/pubsub"
	"net"
)

// Server is a struct that represents a server
type Server struct {
	listener       net.Listener
	port           string
	logger         logger.Logger
	cache          cache.Cache
	peers          []peers.Peer
	messagingQueue pubsub.Queue
}

// NewServer creates a new server
func NewServer(options ServerOptions) server.FamcacheServer {
	return &Server{
		port:           options.Port,
		logger:         options.Logger,
		cache:          options.Cache,
		peers:          make([]peers.Peer, 0),
		messagingQueue: queue.NewPubsubQueue(),
	}
}
