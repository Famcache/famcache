package pubsub

import "famcache/domain/pubsub"

func (q *PubsubQueue) Retry(message pubsub.Message) {
	q.messages = append(q.messages, message)
}
