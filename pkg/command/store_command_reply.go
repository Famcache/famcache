package command

import (
	"fmt"
	"net"
)

const storeMark = "STORE"

func (q *StoreCommand) ReplyOK(conn net.Conn, response string) {
	message := fmt.Sprintf("%s %s OK %s\n", q.ID(), storeMark, response)

	conn.Write([]byte(message))
}

func (q *StoreCommand) ReplySuccess(conn net.Conn) {
	message := fmt.Sprintf("%s %s OK\n", q.ID(), storeMark)

	conn.Write([]byte(message))
}

func (q *StoreCommand) ReplyError(conn net.Conn, response string) {
	message := fmt.Sprintf("%s %s ERROR %s\n", q.ID(), storeMark, response)

	conn.Write([]byte(message))
}
