package actor

import (
	"famcache/domain/command"
	"famcache/domain/connection"
	"fmt"
)

func (actor *Actor) RegisterJob(peer connection.Peer, query command.JobCommand) {
	logger := *actor.logger

	id := actor.jobs.Add(peer.ID(), query.Delay(), query.IsPeriodic())

	logger.Info(fmt.Sprintf("Job registered with id %s", id))

	query.ReplyOK(peer.Conn(), id)
}
