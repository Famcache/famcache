package connection

type Subscription interface {
	Add(topic string)
	Remove(topic string)
	IsSubscribed(topic string) bool
}
