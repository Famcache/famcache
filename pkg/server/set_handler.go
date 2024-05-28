package server

import (
	"famcache/domain/cache"
	"net"
)

// handleSet handles the SET command
// set <key> <value> <ttl>
func (s *Server) handleSet(conn net.Conn, query *Query) {
	err := s.cache.Set(cache.SetOptions{
		Key:   query.Key,
		Value: *query.Value,
		TTL:   query.TTL,
	})

	if err != nil {
		s.logger.Error("Error setting key")

		query.replyError(conn, err.Error())

		return
	}

	s.logger.Info("SET", query.Key, *query.Value)

	query.replySuccess(conn)
}
