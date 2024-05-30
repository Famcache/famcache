package server

import (
	"bufio"
	"famcache/pkg/server/command"
	"fmt"
	"net"
)

func (s *Server) handle(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		request, err := reader.ReadString('\n')

		if err != nil {
			s.logger.Error("Error reading request")
			return
		}

		println("Request: ", request)
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

			switch query.Type {
			case command.CommandGet:
				s.handleGet(conn, query)
			case command.CommandSet:
				s.handleSet(conn, query)
			case command.CommandDelete:
				s.handleDelete(conn, query)
			default:
				println("Invalid command")
			}
		}

		if com.IsPublishCommand() {
			action := com.ToPubsubCommand()

			fmt.Println(action)
		}
	}
}
