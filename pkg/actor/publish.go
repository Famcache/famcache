package actor

import (
	"famcache/domain/command"
	"famcache/domain/connection"
	"famcache/pkg/pubsub"
)

func (actor *Actor) Publish(peer connection.Peer, message command.MessagingCommand) error {
	logger := *actor.logger

	logger.Info("Peer " + peer.ID() + " published topic " + message.Topic())

	for _, p := range *actor.peers.All() {
		if p.ID() == peer.ID() {
			continue
		}

		if !p.Subscriptions().IsSubscribed(message.Topic()) {
			continue
		}

		// Try to publish the message to the peer immediately
		// If the peer is not available, keep the message in the queue
		// and retry later
		err := p.Publish(message.Topic(), message.Data())

		if err != nil {
			// TODO: Retry logic
			actor.messagingQueue.Enqueue(pubsub.NewPubsubMessage(peer.ID(), message.Topic(), message.Data()))
		}
	}

	return nil
}
