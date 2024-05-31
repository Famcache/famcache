package actor

import (
	"famcache/domain/connection"
	"famcache/pkg/server/command"
)

func (actor *Actor) Subscribe(peer connection.Peer, query *command.MessagingCommand) {
	logger := *actor.logger

	logger.Info("Peer" + peer.ID() + " subscribed to the topic: " + query.Topic)

	sub := peer.Subscriptions()

	sub.Add(query.Topic)
}
