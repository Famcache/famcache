package peers

import (
	"net"

	"github.com/google/uuid"
)

type Peer struct {
	id   string
	conn net.Conn
}

func (p *Peer) ID() string {
	return p.id
}

func (p *Peer) Conn() net.Conn {
	return p.conn
}

func NewPeer(conn net.Conn) *Peer {
	id := uuid.NewString()

	return &Peer{id, conn}
}