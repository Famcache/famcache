package server

import (
	"famcache/pkg/server/command"
	"net"
)

func (s *Server) handleDelete(conn net.Conn, query *command.StoreCommand) {
	err := s.cache.Delete(query.Key)

	if err != nil {
		s.logger.Error("Error deleting key")

		query.ReplyError(conn, err.Error())

		return
	}

	s.logger.Info("DELETE", query.Key)

	query.ReplySuccess(conn)
}
