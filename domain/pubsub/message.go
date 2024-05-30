package pubsub

type Message interface {
	GetTopic() string
	GetData() string
}
