package jobs_test

import (
	"famcache/pkg/jobs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJobManager(t *testing.T) {
	jm := jobs.NewJobsManager()

	assert.NotNil(t, jm)
}

func TestAddJob(t *testing.T) {
	jm := jobs.NewJobsManager()

	id := jm.Add("peer-id", 5000, true)

	assert.NotEmpty(t, id)
}

func TestCancelJob(t *testing.T) {
	jm := jobs.NewJobsManager()
	peerId := "peer-id"

	id := jm.Add(peerId, 5000, true)

	jm.Cancel(id)

	_, ok := jm.Get(id)

	assert.False(t, ok)
}
