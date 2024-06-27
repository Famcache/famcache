package server

import (
	"famcache/domain/command"
	"famcache/domain/connection"
)

func (s *Server) handleJobCommand(peer connection.Peer, com command.JobCommand) {
	switch com.Type() {
	case command.CommandJobRegister:
		s.actor.RegisterJob(peer, com)
	case command.CommandJobCancel:
		s.actor.CancelJob(peer, com)
	default:
		println("Invalid command")
	}
}
