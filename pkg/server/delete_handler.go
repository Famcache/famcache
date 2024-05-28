package server

import "net"

func (s *Server) handleDelete(conn net.Conn, query *Query) {
	err := s.cache.Delete(query.Key)

	if err != nil {
		s.logger.Error("Error deleting key")

		query.replyError(conn, err.Error())

		return
	}

	s.logger.Info("DELETE", query.Key)

	query.replySuccess(conn)
}
