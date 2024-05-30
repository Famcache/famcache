package server

import (
	"famcache/pkg/server/command"
	"famcache/pkg/server/peers"
)

func (s *Server) handleUnsubscribe(peer *peers.Peer, query *command.MessagingCommand) {
	s.logger.Info("Peer" + peer.ID() + " unsubscribed from topic " + query.Topic)

	sub := peer.Subscriptions()

	sub.Remove(query.Topic)
}
