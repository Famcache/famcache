package server

import (
	"famcache/pkg/server/command"
	"net"
)

func (s *Server) handlePublish(conn net.Conn, query *command.MessagingCommand) {
	s.logger.Info("PUBLISH", query.Topic)
}
