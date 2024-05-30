package server

import (
	"famcache/pkg/server/command"
	"famcache/pkg/server/peers"
)

func (s *Server) handleSubscribe(peer *peers.Peer, query *command.MessagingCommand) {
	s.logger.Info("Peer" + peer.ID() + " subscribed to the topic: " + query.Topic)

	sub := peer.Subscriptions()

	sub.Add(query.Topic)
}
