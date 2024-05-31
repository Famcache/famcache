package actor

import (
	"famcache/pkg/server/command"
	"net"
)

func (actor *Actor) Get(conn net.Conn, query *command.StoreCommand) {
	sCache := *actor.cache
	logger := *actor.logger

	value, err := sCache.Get(query.Key)

	if err != nil {
		logger.Error("Error getting key")

		query.ReplyError(conn, err.Error())
	}

	logger.Info("GET", query.Key, value)

	query.ReplyOK(conn, value)
}
