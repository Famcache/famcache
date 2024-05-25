package server

import "net"

func (s *Server) handleDelete(conn net.Conn, parts []string) {
	if len(parts) != 2 {
		s.logger.Error("Invalid DELETE request")

		s.replyError(conn, "Invalid DELETE request")
	}

	key := parts[1]

	err := s.cache.Delete(key)

	if err != nil {
		s.logger.Error("Error deleting key")

		s.replyError(conn, err.Error())
	}

	s.logger.Info("DELETE", key)

	s.replySuccess(conn)
}
