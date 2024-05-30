package server

import (
	"famcache/pkg/server/command"
	"net"
)

func (s *Server) handleGet(conn net.Conn, query *command.StoreCommand) {
	value, err := s.cache.Get(query.Key)

	if err != nil {
		s.logger.Error("Error getting key")

		query.ReplyError(conn, err.Error())
	}

	s.logger.Info("GET", query.Key, value)

	query.ReplyOK(conn, value)
}
