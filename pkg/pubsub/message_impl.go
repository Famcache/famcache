package pubsub

import (
	"famcache/domain/command"
	cmd "famcache/pkg/command"
)

func (m *PubsubMessage) ID() string {
	return m.id
}

func (m *PubsubMessage) Topic() string {
	return m.topic
}

func (m *PubsubMessage) Data() string {
	return m.data
}

func (m *PubsubMessage) Recipient() string {
	return m.to
}

func (m *PubsubMessage) CreatedAt() int64 {
	return m.createdAt
}

func (m *PubsubMessage) RetryCount() uint {
	return m.retried
}

func (m *PubsubMessage) IncrementRetryCount() {
	m.retried++
}

func (m *PubsubMessage) ToMessagingCommand() command.MessagingCommand {
	query := cmd.PublishCommand("internal", m.Topic(), m.Data())

	println("Query", query)
	return cmd.NewPubsubCommand(command.CommandPublish, query)
}
