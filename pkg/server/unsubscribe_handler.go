package server

import (
	"famcache/pkg/server/command"
	"net"
)

func (s *Server) handleUnsubscribe(conn net.Conn, query *command.MessagingCommand) {
}
