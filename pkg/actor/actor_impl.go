package actor

import "famcache/pkg/command"

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

func (actor *Actor) ListenJobs() {
	for {
		job := <-actor.jobs.Chan()
		peerId := job.PeerId()

		peer := actor.peers.GetById(peerId)

		if peer == nil {
			continue
		}

		command := command.JobExecuteCommand(job.ID())

		peer.Conn().Write([]byte(command))
	}
}

func (actor *Actor) Start() {
	go actor.ListenFailedMessages()
	go actor.ListenJobs()
}
