package server

import (
	"famcache/pkg/cache"
	"famcache/pkg/config"
	"famcache/pkg/logger"
	"testing"
)

func TestNewServer(t *testing.T) {
	logger := logger.NewLogger()
	cache := cache.NewCache()
	config := config.NewConfig()

	s := NewServer(ServerOptions{
		Port:   config.GetPort(),
		Logger: logger,
		Cache:  cache,
	})

	if s == nil {
		t.Error("Expected server instance, got nil")
	}

	if _, ok := s.(*Server); !ok {
		t.Error("Expected Server instance, got different type")
	}

	t.Log("NewServer() test passed")
}
