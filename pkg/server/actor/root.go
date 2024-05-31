package actor

import (
	"famcache/domain/cache"
	"famcache/domain/connection"
	"famcache/domain/logger"
	"famcache/domain/pubsub"
	conn "famcache/pkg/connection"
	queue "famcache/pkg/pubsub"
)

type Actor struct {
	logger         *logger.Logger
	cache          *cache.Cache
	messagingQueue pubsub.Queue
	peers          connection.PeersManager
}

func (a *Actor) Peers() *connection.PeersManager {
	return &a.peers
}

func NewActor(logger *logger.Logger, cache *cache.Cache) Actor {
	messagingQueue := queue.NewPubsubQueue()
	peers := conn.NewPeersManager()

	return Actor{
		logger,
		cache,
		messagingQueue,
		peers,
	}
}
