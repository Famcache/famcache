package cache

import (
	"famcache/domain"
	"strconv"
	"strings"
	"time"
)

type StoredValue struct {
	Value     string
	TTL       *int64
	CreatedAt int64
}

func NewStoredValue(value string, ttl *string) (*StoredValue, error) {
	var ttlInt64 *int64

	if ttl != nil {
		ttlInt64 = new(int64)
		ttlConverted, err := strconv.ParseInt(strings.TrimSpace(*ttl), 10, 64)

		if err != nil {
			return nil, domain.ErrInvalidTTL
		}

		*ttlInt64 = ttlConverted
	}

	return &StoredValue{
		Value:     value,
		TTL:       ttlInt64,
		CreatedAt: time.Now().Unix(),
	}, nil
}
