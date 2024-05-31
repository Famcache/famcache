package connection

import "net"

type Peer interface {
	ID() string
	Subscriptions() Subscription
	Publish(topic, data string) error
	Conn() net.Conn
}
