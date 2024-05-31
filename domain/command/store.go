package command

import "net"

type StoreCommand interface {
	ID() string
	Type() CommandType
	Key() string
	Value() *string
	TTL() *uint64

	ReplyError(conn net.Conn, response string)
	ReplySuccess(conn net.Conn)
	ReplyOK(conn net.Conn, response string)
}
