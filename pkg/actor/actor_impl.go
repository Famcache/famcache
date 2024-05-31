package actor

func (actor *Actor) ListenFailedMessages() {
	for {
		<-actor.queueTicker.C

		for _, message := range actor.messagingQueue.Batch() {
			peer := actor.peers.GetById(message.Recipient())

			if peer == nil {
				actor.messagingQueue.Remove(message.ID())
				continue
			}

			err := actor.Publish(peer, message.ToMessagingCommand())

			if err == nil {
				actor.messagingQueue.Remove(message.ID())
			} else {
				message.IncrementRetryCount()
			}
		}
	}
}
