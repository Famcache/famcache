package server

import (
	"famcache/domain/connection"
	"famcache/pkg/server/command"
)

func (s *Server) handleMessagingCommand(peer connection.Peer, com *command.MessagingCommand) {
	switch com.Type {
	case command.CommandPublish:
		s.actor.Publish(peer, com)
	case command.CommandSubscribe:
		s.actor.Subscribe(peer, com)
	case command.CommandUnsubscribe:
		s.actor.Unsubscribe(peer, com)
	default:
		println("Invalid command")
	}
}
