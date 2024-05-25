package server

import (
	"famcache/domain/cache"
	"famcache/domain/logger"
)

type ServerOptions struct {
	Port   string
	Logger logger.Logger
	Cache  cache.Cache
}
