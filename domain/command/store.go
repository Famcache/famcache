package command

type StoreCommand interface {
	ID() string
	Type() CommandType
	Key() string
	Value() *string
	TTL() *uint64

	Reply
}
