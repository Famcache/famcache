package pubsub

import "famcache/domain/pubsub"

type PubsubMessage struct {
	to      string
	topic   string
	data    string
	retried uint
}

func NewPubsubMessage(to, topic, data string) pubsub.Message {
	var retried uint = 0

	return &PubsubMessage{
		to,
		topic,
		data,
		retried,
	}
}
