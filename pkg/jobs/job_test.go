package jobs_test

import (
	"famcache/pkg/jobs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJob(t *testing.T) {
	delay := uint64(5000)
	isPeriodic := true

	job := jobs.NewJob("peer-id", delay, isPeriodic)

	assert.Equal(t, delay, job.Delay())
	assert.Equal(t, isPeriodic, job.IsPeriodic())
}

func TestJobMethods(t *testing.T) {
	delay := uint64(5000)
	isPeriodic := true
	peerId := "peer-id"

	job := jobs.NewJob(peerId, delay, isPeriodic)

	assert.NotEmpty(t, job.ID())
	assert.Equal(t, delay, job.Delay())
	assert.Equal(t, isPeriodic, job.IsPeriodic())
	assert.Equal(t, peerId, job.PeerId())
}
