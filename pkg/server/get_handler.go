package server

import "net"

func (s *Server) handleGet(conn net.Conn, query *Query) {
	value, err := s.cache.Get(query.Key)

	if err != nil {
		s.logger.Error("Error getting key")

		query.replyError(conn, err.Error())
	}

	s.logger.Info("GET", query.Key, value)

	query.replyOK(conn, value)
}
