package command

import "net"

type Reply interface {
	ReplyError(conn net.Conn, response string)
	ReplySuccess(conn net.Conn)
	ReplyOK(conn net.Conn, response string)
}
