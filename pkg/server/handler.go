package server

import (
	"bufio"
	"famcache/domain/connection"
	"famcache/pkg/server/command"
	"io"
)

func (s *Server) handle(peer connection.Peer) {
	reader := bufio.NewReader(peer.Conn())

	for {
		request, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				(*s.actor.Peers()).Remove(peer)
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
