package actor

import (
	"famcache/domain/command"
	"famcache/domain/connection"
)

func (a *Actor) CancelJob(peer connection.Peer, query command.JobCommand) {
	logger := *a.logger

	a.jobs.Cancel(query.JobId())

	logger.Info("Job canceled", "id", query.JobId())

	query.ReplyOK(peer.Conn(), "OK")
}
