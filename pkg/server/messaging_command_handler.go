package server

import (
	"famcache/pkg/server/command"
	"famcache/pkg/server/peers"
)

func (s *Server) handleMessagingCommand(peer *peers.Peer, com *command.MessagingCommand) {
	switch com.Type {
	case command.CommandPublish:
		s.handlePublish(peer, com)
	case command.CommandSubscribe:
		s.handleSubscribe(peer, com)
	case command.CommandUnsubscribe:
		s.handleUnsubscribe(peer, com)
	default:
		println("Invalid command")
	}
}