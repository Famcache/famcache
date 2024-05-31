package command

import "net"

func (q *StoreCommand) ReplyOK(conn net.Conn, response string) {
	message := q.ID() + " OK " + response + "\n"

	conn.Write([]byte(message))
}

func (q *StoreCommand) ReplySuccess(conn net.Conn) {
	message := q.ID() + " OK" + "\n"

	conn.Write([]byte(message))
}

func (q *StoreCommand) ReplyError(conn net.Conn, response string) {
	message := q.ID() + " ERROR: " + response + "\n"

	conn.Write([]byte(message))
}
