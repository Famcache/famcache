package command

import (
	"fmt"
	"net"
)

const jobMark = "JOB"

func (q *JobCommand) ReplyOK(conn net.Conn, response string) {
	message := fmt.Sprintf("%s %s OK %s\n", q.ID(), jobMark, response)

	conn.Write([]byte(message))
}

func (q *JobCommand) ReplySuccess(conn net.Conn) {
	message := fmt.Sprintf("%s %s OK\n", q.ID(), jobMark)

	conn.Write([]byte(message))
}

func (q *JobCommand) ReplyError(conn net.Conn, response string) {
	message := fmt.Sprintf("%s %s ERROR %s\n", q.ID(), jobMark, response)

	conn.Write([]byte(message))
}
