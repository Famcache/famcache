package pubsub

import "famcache/domain/command"

type Message interface {
	ID() string
	Topic() string
	Data() string
	Recipient() string
	CreatedAt() int64
	RetryCount() uint
	IncrementRetryCount()
	ToMessagingCommand() command.MessagingCommand
}
