package server

type FamcacheServer interface {
	Start() error
	Stop() error
}
