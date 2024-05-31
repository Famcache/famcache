package actor

import (
	"famcache/domain/connection"
	"famcache/pkg/server/command"
)

func (actor *Actor) Unsubscribe(peer connection.Peer, query *command.MessagingCommand) {
	logger := *actor.logger

	logger.Info("Peer" + peer.ID() + " unsubscribed from topic " + query.Topic)

	sub := peer.Subscriptions()

	sub.Remove(query.Topic)
}
