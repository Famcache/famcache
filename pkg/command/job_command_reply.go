package command

import "net"

func (q *JobCommand) ReplyOK(conn net.Conn, response string) {
	message := q.ID() + " OK " + response + "\n"

	conn.Write([]byte(message))
}

func (q *JobCommand) ReplySuccess(conn net.Conn) {
	message := q.ID() + " OK" + "\n"

	conn.Write([]byte(message))
}

func (q *JobCommand) ReplyError(conn net.Conn, response string) {
	message := q.ID() + " ERROR: " + response + "\n"

	conn.Write([]byte(message))
}
