package command

func (sc *StoreCommand) ID() string {
	return sc.id
}

func (sc *StoreCommand) Type() string {
	return sc.cType
}

func (sc *StoreCommand) Key() string {
	return sc.key
}

func (sc *StoreCommand) Value() *string {
	return sc.value
}

func (sc *StoreCommand) TTL() *uint64 {
	return sc.ttl
}
