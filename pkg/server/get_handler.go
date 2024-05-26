package server

import "net"

func (s *Server) handleGet(conn net.Conn, parts []string) {
	if len(parts) != 2 {
		s.logger.Error("Invalid GET request")

		s.replyError(conn, "Invalid GET request")

		return
	}

	key := parts[1]
	value, err := s.cache.Get(key)

	if err != nil {
		s.logger.Error("Error getting key")

		s.replyError(conn, err.Error())
	}

	s.logger.Info("GET", key, value)

	s.replyOK(conn, value)
}
