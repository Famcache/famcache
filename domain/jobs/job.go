package jobs

type Job interface {
	ID() string
	Delay() uint64
	IsPeriodic() bool
	PeerId() string
}
