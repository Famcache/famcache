package main

import (
	"famcache/pkg/cache"
	"famcache/pkg/config"
	"famcache/pkg/logger"
	"famcache/pkg/server"
	"log"
)

func main() {
	config := config.NewConfig()

	cache := cache.NewCache()
	options := server.ServerOptions{
		Port:   config.Port,
		Logger: logger.NewLogger(),
		Cache:  cache,
	}

	server := server.NewServer(options)

	err := server.Start()

	if err != nil {
		log.Fatal(err)
	}
}
