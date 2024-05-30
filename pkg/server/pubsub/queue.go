package pubsub

import "famcache/domain/pubsub"

type PubsubQueue struct {
	messages []pubsub.Message
}

func NewPubsubQueue() pubsub.Queue {
	return &PubsubQueue{
		messages: []pubsub.Message{},
	}
}
