package pubsub

type Queue interface {
	Retry(message Message)
}
