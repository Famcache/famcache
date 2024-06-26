package jobs

func (j *Job) ID() string {
	return j.id
}

func (j *Job) Delay() uint64 {
	return j.delay
}

func (j *Job) IsPeriodic() bool {
	return j.isPeriodic
}

func (j *Job) PeerId() string {
	return j.peerId
}
