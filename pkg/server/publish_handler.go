package server

import (
	"famcache/pkg/server/command"
	"famcache/pkg/server/peers"
	"famcache/pkg/server/pubsub"
)

func (s *Server) handlePublish(peer *peers.Peer, message *command.MessagingCommand) {
	s.logger.Info("Peer " + peer.ID() + " published topic " + message.Topic)

	for _, p := range *s.peers.All() {
		if p.ID() == peer.ID() {
			continue
		}

		// Try to publish the message to the peer immediately
		// If the peer is not available, keep the message in the queue
		// and retry later
		err := p.Publish(message.Topic, message.Data)

		if err != nil {
			// TODO: Retry logic
			s.messagingQueue.Retry(pubsub.NewPubsubMessage(peer.ID(), message.Topic, message.Data))
		}
	}
}
