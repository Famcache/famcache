package command

import (
	"famcache/domain/command"
	"strconv"
	"strings"
)

type StoreCommand struct {
	// Common fields
	cType command.CommandType
	id    string

	// Get and Delete fields
	key string

	// Set fields
	value *string
	ttl   *uint64
}

func NewStoreCommand(commandType command.CommandType, query string) command.StoreCommand {
	parts := strings.Fields(strings.TrimSpace(query))

	if len(parts) < 3 {
		return nil
	}

	queryId := parts[0]

	key := parts[2]
	var ttl *uint64 = nil
	var value *string = nil

	if commandType == command.CommandSet {
		value = &parts[3]

		if len(parts) == 5 {
			convTtl, err := strconv.ParseUint(parts[4], 10, 64)
			if err != nil {
				return nil
			}
			ttl = &convTtl
		}
	}

	return &StoreCommand{
		cType: commandType,
		id:    queryId,
		key:   key,
		value: value,
		ttl:   ttl,
	}
}
