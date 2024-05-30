package pubsub

func (m *PubsubMessage) GetTopic() string {
	return m.topic
}

func (m *PubsubMessage) GetData() string {
	return m.data
}
