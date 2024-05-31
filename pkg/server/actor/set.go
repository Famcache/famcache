package actor

import (
	"famcache/domain/cache"
	"famcache/pkg/server/command"
	"net"
)

func (actor *Actor) Set(conn net.Conn, query *command.StoreCommand) {
	sCache := *actor.cache
	logger := *actor.logger

	err := sCache.Set(cache.SetOptions{
		Key:   query.Key,
		Value: *query.Value,
		TTL:   query.TTL,
	})

	if err != nil {
		logger.Error("Error setting key")

		query.ReplyError(conn, err.Error())

		return
	}

	logger.Info("SET", query.Key, *query.Value)

	query.ReplySuccess(conn)
}
