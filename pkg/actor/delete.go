package actor

import (
	"famcache/domain/command"
	"net"
)

func (actor *Actor) Delete(conn net.Conn, query command.StoreCommand) {
	sCache := *actor.cache
	logger := *actor.logger

	err := sCache.Delete(query.Key())

	if err != nil {
		logger.Error("Error deleting key")

		query.ReplyError(conn, err.Error())

		return
	}

	logger.Info("DELETE", query.Key())

	query.ReplySuccess(conn)
}
