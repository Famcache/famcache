package connection

type Subscriptions struct {
	inner map[string]bool
}

func (s *Subscriptions) Add(topic string) {
	s.inner[topic] = true
}

func (s *Subscriptions) Remove(topic string) {
	delete(s.inner, topic)
}

func (s *Subscriptions) IsSubscribed(topic string) bool {
	return s.inner[topic]
}

func NewSubscription() Subscriptions {
	return Subscriptions{
		inner: make(map[string]bool),
	}
}
