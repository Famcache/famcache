package actor

import (
	"famcache/domain/command"
	"famcache/domain/connection"
)

func (actor *Actor) Subscribe(peer connection.Peer, query command.MessagingCommand) {
	logger := *actor.logger

	logger.Info("Peer" + peer.ID() + " subscribed to the topic: " + query.Topic())

	sub := peer.Subscriptions()

	sub.Add(query.Topic())
}
