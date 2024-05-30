package server

import (
	"famcache/pkg/server/command"
	"famcache/pkg/server/peers"
)

func (s *Server) handleStoreCommand(peer *peers.Peer, com *command.StoreCommand) {
	switch com.Type {
	case command.CommandGet:
		s.handleGet(peer.Conn(), com)
	case command.CommandSet:
		s.handleSet(peer.Conn(), com)
	case command.CommandDelete:
		s.handleDelete(peer.Conn(), com)
	default:
		println("Invalid command")
	}
}
