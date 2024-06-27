package jobs

import (
	job "famcache/domain/jobs"

	"github.com/google/uuid"
)

type Job struct {
	id         string
	delay      uint64
	isPeriodic bool
	peerId     string
}

func NewJob(peerId string, delay uint64, isPeriodic bool) job.Job {
	id := uuid.New().String()

	return &Job{
		id:         id,
		delay:      delay,
		isPeriodic: isPeriodic,
		peerId:     peerId,
	}
}
