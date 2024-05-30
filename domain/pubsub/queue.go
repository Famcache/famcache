package pubsub

type Queue interface {
	Publish(message Message)
}
