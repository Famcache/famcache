package jobs

import (
	job "famcache/domain/jobs"
	mJob "famcache/domain/jobs"
	"time"
)

func (jm *JobsManager) StartJob(job mJob.Job) {
	time.Sleep(time.Duration(job.Delay()) * time.Millisecond)

	jm.triggerChan <- job

	if !job.IsPeriodic() {
		jm.Cancel(job.ID())
		return
	}

	jm.StartJob(job)
}

func (jm *JobsManager) Add(peerId string, delay uint64, isPeriodic bool) string {
	job := NewJob(peerId, delay, isPeriodic)

	jm.jobs[job.ID()] = job

	go jm.StartJob(job)

	return job.ID()
}

func (jm *JobsManager) Get(id string) (job.Job, bool) {
	job, ok := jm.jobs[id]

	return job, ok
}

func (jm *JobsManager) Cancel(id string) {
	delete(jm.jobs, id)
}

func (jm *JobsManager) Chan() <-chan job.Job {
	return jm.triggerChan
}

func (jm *JobsManager) Jobs() *map[string]mJob.Job {
	return &jm.jobs
}
