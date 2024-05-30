package server

import (
	"strconv"
	"strings"
)

type QueryType = string

const (
	QueryTypeGet    QueryType = "GET"
	QueryTypeSet    QueryType = "SET"
	QueryTypeDelete QueryType = "DELETE"
)

type Query struct {
	// Common fields
	Type QueryType
	ID   string

	// Get and Delete fields
	Key string

	// Set fields
	Value *string
	TTL   *int64
}

var strToQueryType = map[string]QueryType{
	"GET":    QueryTypeGet,
	"SET":    QueryTypeSet,
	"DELETE": QueryTypeDelete,
}

func determineQueryType(queryType string) (QueryType, bool) {
	t, ok := strToQueryType[strings.TrimSpace(strings.ToUpper((queryType)))]
	return t, ok
}

func NewQuery(query string) *Query {
	parts := strings.Fields(strings.TrimSpace(query))

	println(query)
	if len(parts) < 3 {
		return nil
	}

	queryId := parts[0]
	queryType, ok := determineQueryType(parts[1])

	if !ok {
		return nil
	}

	key := parts[2]
	var ttl *int64 = nil
	var value *string = nil

	if queryType == QueryTypeSet {
		value = &parts[3]

		if len(parts) == 5 {
			convTtl, err := strconv.ParseInt(parts[4], 10, 64)
			if err != nil {
				return nil
			}
			ttl = &convTtl
		}
	}

	return &Query{
		Type:  queryType,
		ID:    queryId,
		Key:   key,
		Value: value,
		TTL:   ttl,
	}
}
