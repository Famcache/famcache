package server

import (
	"famcache/domain/cache"
	"famcache/pkg/server/command"
	"net"
)

// handleSet handles the SET command
// set <key> <value> <ttl>
func (s *Server) handleSet(conn net.Conn, query *command.StoreCommand) {
	err := s.cache.Set(cache.SetOptions{
		Key:   query.Key,
		Value: *query.Value,
		TTL:   query.TTL,
	})

	if err != nil {
		s.logger.Error("Error setting key")

		query.ReplyError(conn, err.Error())

		return
	}

	s.logger.Info("SET", query.Key, *query.Value)

	query.ReplySuccess(conn)
}
