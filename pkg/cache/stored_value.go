package cache

import (
	"time"
)

type StoredValue struct {
	Value     string
	TTL       *uint64
	CreatedAt int64
}

func NewStoredValue(value string, ttl *uint64) (*StoredValue, error) {
	return &StoredValue{
		Value:     value,
		TTL:       ttl,
		CreatedAt: time.Now().Unix(),
	}, nil
}
