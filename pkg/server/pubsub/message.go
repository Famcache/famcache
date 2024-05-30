package pubsub

import "famcache/domain/pubsub"

type PubsubMessage struct {
	to    string
	topic string
	data  string
}

func NewPubsubMessage(to, topic, data string) pubsub.Message {
	return &PubsubMessage{
		to,
		topic,
		data,
	}
}
