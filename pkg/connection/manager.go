package connection

import (
	"famcache/domain/connection"
	"net"
)

type PeersManager struct {
	peers map[string]connection.Peer
}

func (pm *PeersManager) Add(conn net.Conn) connection.Peer {
	peer := NewPeer(conn)
	pm.peers[peer.ID()] = peer

	return peer
}

func (pm *PeersManager) Remove(peer connection.Peer) {
	delete(pm.peers, peer.ID())
}

func (pm *PeersManager) All() *map[string]connection.Peer {
	return &pm.peers
}

func (pm *PeersManager) GetById(id string) connection.Peer {
	return pm.peers[id]
}

func NewPeersManager() connection.PeersManager {
	return &PeersManager{
		peers: make(map[string]connection.Peer),
	}
}
