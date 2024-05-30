package pubsub

import "famcache/domain/pubsub"

type PubsubMessage struct {
	topic string
	data  string
}

func NewPubsubMessage(topic string, data string) pubsub.Message {
	return &PubsubMessage{
		topic,
		data,
	}
}
