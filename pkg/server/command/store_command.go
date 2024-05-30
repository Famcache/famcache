package command

import (
	"strconv"
	"strings"
)

type StoreCommand struct {
	// Common fields
	Type CommandType
	ID   string

	// Get and Delete fields
	Key string

	// Set fields
	Value *string
	TTL   *int64
}

func NewStoreCommand(commandType CommandType, query string) *StoreCommand {
	parts := strings.Fields(strings.TrimSpace(query))

	if len(parts) < 3 {
		return nil
	}

	queryId := parts[0]

	key := parts[2]
	var ttl *int64 = nil
	var value *string = nil

	if commandType == CommandSet {
		value = &parts[3]

		if len(parts) == 5 {
			convTtl, err := strconv.ParseInt(parts[4], 10, 64)
			if err != nil {
				return nil
			}
			ttl = &convTtl
		}
	}

	return &StoreCommand{
		Type:  commandType,
		ID:    queryId,
		Key:   key,
		Value: value,
		TTL:   ttl,
	}
}
