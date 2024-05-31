package pubsub

import (
	"famcache/domain/pubsub"
	"time"

	"github.com/google/uuid"
)

type PubsubMessage struct {
	// UUID
	id string

	// Peer ID
	to string

	// Message topic
	topic string

	// Message data
	data string

	// Number of times the message has been retried
	retried uint

	// Message creation time
	createdAt int64
}

func NewPubsubMessage(to, topic, data string) pubsub.Message {
	var retried uint = 0
	createdAt := time.Now().Unix()
	id := uuid.New().String()

	return &PubsubMessage{
		id,
		to,
		topic,
		data,
		retried,
		createdAt,
	}
}
