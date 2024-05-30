package command

import "strings"

type PubsubCommand struct {
	ID    string
	Type  CommandType
	Topic string
	Data  string
}

func NewPubsubCommand(commandType CommandType, query string) *PubsubCommand {
	parts := strings.Fields(strings.TrimSpace(query))

	if len(parts) < 3 {
		return nil
	}

	queryId := parts[0]

	topic := parts[2]
	var data string

	if commandType == CommandPublish {
		if len(parts) < 4 {
			data = ""
		} else {
			data = parts[3]
		}
	}

	return &PubsubCommand{
		ID:    queryId,
		Topic: topic,
		Type:  commandType,
		Data:  data,
	}
}
