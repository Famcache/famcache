package peers

import (
	"net"

	"github.com/google/uuid"
)

type Peer struct {
	id            string
	conn          net.Conn
	subscriptions Subscriptions
}

func (p *Peer) ID() string {
	return p.id
}

func (p *Peer) Subscriptions() Subscriptions {
	return p.subscriptions
}

func (p *Peer) Publish(topic, data string) error {
	msg := "MESSAGE " + topic + " " + data + "\n"

	_, err := p.conn.Write([]byte(msg))

	return err
}

func (p *Peer) Conn() net.Conn {
	return p.conn
}

func NewPeer(conn net.Conn) *Peer {
	id := uuid.NewString()
	subscriptions := NewSubscription()

	return &Peer{id, conn, subscriptions}
}
