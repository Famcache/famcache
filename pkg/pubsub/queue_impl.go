package pubsub

import (
	"famcache/domain/pubsub"
	"time"
)

func (q *PubsubQueue) Enqueue(message pubsub.Message) {
	q.messages = append(q.messages, message)
}

func (q *PubsubQueue) Batch() []pubsub.Message {
	// Filter out messages that needs to be retried
	// Message should be retried if
	// - {currentTime} - message.createdAt >= message.retried * 5 seconds

	var messages []pubsub.Message = []pubsub.Message{}

	for _, message := range q.messages {
		if message.RetryCount() == 0 {
			messages = append(messages, message)
		}

		elapsedTime := time.Now().Unix() - message.CreatedAt()

		if elapsedTime >= int64(message.RetryCount()*5) {
			messages = append(messages, message)
		}
	}

	return messages
}

func (q *PubsubQueue) Remove(messageId string) {
	for i, message := range q.messages {
		if message.ID() == messageId {
			q.messages = append(q.messages[:i], q.messages[i+1:]...)
			break
		}
	}
}
