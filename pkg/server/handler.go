package server

import (
	"bufio"
	"famcache/pkg/server/command"
	"famcache/pkg/server/peers"
	"io"
)

func (s *Server) handle(peer *peers.Peer) {
	reader := bufio.NewReader(peer.Conn())

	for {
		request, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				s.peers.Remove(peer)
				return
			}

			s.logger.Error("Error reading request")
			return
		}

		com, err := command.NewCommand(request)

		if err != nil {
			s.logger.Error("Error parsing request")
			continue
		}

		if com == nil {
			s.logger.Error("Empty request")
			continue
		}

		if com.IsStoreCommand() {
			query := com.ToStoreCommand()

			s.handleStoreCommand(peer, query)
		}

		if com.IsMessagingCommand() {
			action := com.ToPubsubCommand()

			s.handleMessagingCommand(peer, action)
		}
	}
}
