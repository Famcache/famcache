package command

type MessagingCommand interface {
	ID() string
	Type() CommandType
	Topic() string
	Data() string
}
