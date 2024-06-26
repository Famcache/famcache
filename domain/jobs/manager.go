package jobs

type JobsManager interface {
	Add(peerId string, delay uint64, isPeriodic bool) string
	Get(id string) (Job, bool)
	Cancel(id string)
	Chan() <-chan Job
}
