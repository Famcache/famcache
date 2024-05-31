package server

import (
	"famcache/domain/command"
	"famcache/domain/connection"
)

func (s *Server) handleStoreCommand(peer connection.Peer, com command.StoreCommand) {
	switch com.Type() {
	case command.CommandGet:
		s.actor.Get(peer.Conn(), com)
	case command.CommandSet:
		s.actor.Set(peer.Conn(), com)
	case command.CommandDelete:
		s.actor.Delete(peer.Conn(), com)
	default:
		println("Invalid command")
	}
}
