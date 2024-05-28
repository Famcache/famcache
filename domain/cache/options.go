package cache

// GetOptions is a struct that contains the options for the Get command
type SetOptions struct {
	Key   string
	Value string
	TTL   *int64
}
