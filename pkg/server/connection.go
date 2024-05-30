package server

import (
	"famcache/pkg/server/peers"
	"net"
)

func (s *Server) AddPeer(conn *net.Conn) string {
	peer := peers.NewPeer(*conn)
	s.peers = append(s.peers, *peer)

	return peer.ID()
}
