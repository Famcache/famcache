package command

import (
	"famcache/domain"
	"famcache/domain/command"
	"strings"
)

var strToCommandType = map[string]command.CommandType{
	"GET":         command.CommandGet,
	"SET":         command.CommandSet,
	"DELETE":      command.CommandDelete,
	"PUBLISH":     command.CommandPublish,
	"SUBSCRIBE":   command.CommandSubscribe,
	"UNSUBSCRIBE": command.CommandUnsubscribe,
}

type AbstractCommand struct {
	cType command.CommandType
	query string
}

func (c *AbstractCommand) ToStoreCommand() command.StoreCommand {
	return NewStoreCommand(c.cType, c.query)
}

func (c *AbstractCommand) ToPubsubCommand() command.MessagingCommand {
	return NewPubsubCommand(c.cType, c.query)
}

func (c *AbstractCommand) IsStoreCommand() bool {
	return c.cType == command.CommandSet || c.cType == command.CommandGet || c.cType == command.CommandDelete
}

func (c *AbstractCommand) IsMessagingCommand() bool {
	return c.cType == command.CommandPublish || c.cType == command.CommandSubscribe || c.cType == command.CommandUnsubscribe
}

func determineCommandType(query string) (command.CommandType, bool) {
	parts := strings.Fields(strings.TrimSpace(query))

	if len(parts) < 2 {
		return command.CommandGet, false
	}

	t, ok := strToCommandType[strings.TrimSpace(strings.ToUpper((parts[1])))]
	return t, ok
}

func NewCommand(query string) (*AbstractCommand, error) {
	commandType, ok := determineCommandType(query)

	if !ok {
		return nil, domain.ErrUnableToProcessRequest
	}

	return &AbstractCommand{
		cType: commandType,
		query: query,
	}, nil
}
