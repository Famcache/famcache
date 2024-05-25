package server

import (
	"bufio"
	"famcache/domain/server"
	"net"
	"strings"
)

func (s *Server) handle(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		request, err := reader.ReadString('\n')

		if err != nil {
			s.logger.Error("Error reading request")
			return
		}

		parts := strings.Fields(strings.TrimSpace(request))

		if len(parts) == 0 {
			s.logger.Error("Empty request")
			continue
		}

		switch parts[0] {
		case server.CommandGet:
			s.handleGet(conn, parts)
		case server.CommandSet:
			s.handleSet(conn, parts)
		case server.CommandDelete:
			s.handleDelete(conn, parts)
		default:
			println("Invalid command")
		}
	}
}
