package command

type JobCommand interface {
	ID() string
	Type() CommandType
	Delay() uint64
	IsPeriodic() bool
	JobId() string

	Reply
}
