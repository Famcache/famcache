package connection

import "net"

type PeersManager interface {
	Add(conn net.Conn) Peer
	Remove(peer Peer)
	All() *map[string]Peer
	GetById(id string) Peer
}
