package server

import "net"

func (s *Server) replyOK(conn net.Conn, response string) {
	message := "OK: " + response

	conn.Write([]byte(message + "\n"))
}

func (s *Server) replySuccess(conn net.Conn) {
	conn.Write([]byte("OK" + "\n"))
}

func (s *Server) replyError(conn net.Conn, response string) {
	message := "ERROR: " + response

	conn.Write([]byte(message + "\n"))
}
