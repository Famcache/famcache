package command

import (
	"famcache/domain"
	"strings"
)

var strToCommandType = map[string]CommandType{
	"GET":       CommandGet,
	"SET":       CommandSet,
	"DELETE":    CommandDelete,
	"PUBLISH":   CommandPublish,
	"SUBSCRIBE": CommandSubscribe,
}

type AbstractCommand struct {
	cType CommandType
	query string
}

func (c *AbstractCommand) ToStoreCommand() *StoreCommand {
	return NewStoreCommand(c.cType, c.query)
}

func (c *AbstractCommand) ToPubsubCommand() *PubsubCommand {
	return NewPubsubCommand(c.cType, c.query)
}

func (c *AbstractCommand) IsStoreCommand() bool {
	return c.cType == CommandSet || c.cType == CommandGet || c.cType == CommandDelete
}

func (c *AbstractCommand) IsPublishCommand() bool {
	return c.cType == CommandPublish || c.cType == CommandSubscribe
}

func determineCommandType(query string) (CommandType, bool) {
	parts := strings.Fields(strings.TrimSpace(query))

	if len(parts) < 2 {
		return CommandGet, false
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
