package jobs

import "famcache/domain/jobs"

type JobsManager struct {
	jobs        map[string]jobs.Job
	triggerChan chan jobs.Job
}

func NewJobsManager() jobs.JobsManager {
	return &JobsManager{
		jobs:        make(map[string]jobs.Job),
		triggerChan: make(chan jobs.Job),
	}
}
