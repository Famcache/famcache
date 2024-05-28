package server

import (
	"bufio"
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

		query := NewQuery(request)

		if query == nil {
			s.logger.Error("Empty request")
			continue
		}

		switch query.Type {
		case QueryTypeGet:
			s.handleGet(conn, query)
		case QueryTypeSet:
			s.handleSet(conn, query)
		case QueryTypeDelete:
			s.handleDelete(conn, query)
		default:
			println("Invalid command")
		}
	}
}
