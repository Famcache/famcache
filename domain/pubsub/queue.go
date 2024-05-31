package pubsub

type Queue interface {
	Enqueue(message Message)
	Remove(messageId string)
	Batch() []Message
}
