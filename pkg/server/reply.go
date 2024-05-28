package server

import "net"

func (q *Query) replyOK(conn net.Conn, response string) {
	message := q.ID + " OK " + response + "\n"

	conn.Write([]byte(message))
}

func (q *Query) replySuccess(conn net.Conn) {
	message := q.ID + " OK" + "\n"

	conn.Write([]byte(message))
}

func (q *Query) replyError(conn net.Conn, response string) {
	message := q.ID + " ERROR: " + response + "\n"

	conn.Write([]byte(message))
}
