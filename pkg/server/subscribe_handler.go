package server

import (
	"famcache/pkg/server/command"
	"net"
)

func (s *Server) handleSubscribe(conn net.Conn, query *command.MessagingCommand) {
	s.logger.Info("SUBSCRIBE", query.Topic)
}
