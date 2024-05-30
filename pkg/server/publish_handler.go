package server

import (
	"famcache/pkg/server/command"
	"famcache/pkg/server/peers"
)

func (s *Server) handlePublish(peer *peers.Peer, message *command.MessagingCommand) {
	s.logger.Info("Peer " + peer.ID() + " published topic " + message.Topic)

	for _, p := range s.peers {
		if p.ID() == peer.ID() {
			continue
		}

		p.Publish(message.Topic, message.Data)
	}
}
