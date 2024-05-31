package command

import (
	"famcache/domain/command"
	"strings"
)

type MessagingCommand struct {
	id    string
	cType command.CommandType
	topic string
	data  string
}

func NewPubsubCommand(commandType command.CommandType, query string) command.MessagingCommand {
	parts := strings.Fields(strings.TrimSpace(query))

	if len(parts) < 3 {
		return nil
	}

	queryId := parts[0]

	topic := parts[2]
	var data string

	if commandType == command.CommandPublish {
		if len(parts) < 4 {
			data = ""
		} else {
			data = parts[3]
		}
	}

	return &MessagingCommand{
		id:    queryId,
		topic: topic,
		cType: commandType,
		data:  data,
	}
}
