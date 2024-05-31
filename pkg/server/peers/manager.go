package peers

import "net"

type PeersManager struct {
	peers map[string]Peer
}

func (pm *PeersManager) Add(conn net.Conn) *Peer {
	peer := NewPeer(conn)
	pm.peers[peer.ID()] = *peer

	return peer
}

func (pm *PeersManager) Remove(peer *Peer) {
	delete(pm.peers, peer.ID())
}

func (pm *PeersManager) All() *map[string]Peer {
	return &pm.peers
}

func (pm *PeersManager) GetById(id string) Peer {
	return pm.peers[id]
}

func NewPeersManager() PeersManager {
	return PeersManager{
		peers: make(map[string]Peer),
	}
}
