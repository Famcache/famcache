package pubsub

import "famcache/domain/pubsub"

func (q *PubsubQueue) Publish(message pubsub.Message) {
	q.messages = append(q.messages, message)
}
