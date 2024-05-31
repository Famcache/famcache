package command

import "famcache/domain/command"

func (m *MessagingCommand) ID() string {
	return m.id
}

func (m *MessagingCommand) Type() command.CommandType {
	return m.cType
}

func (m *MessagingCommand) Topic() string {
	return m.topic
}

func (m *MessagingCommand) Data() string {
	return m.data
}
