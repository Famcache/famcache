package server

import (
	"famcache/domain"
	"fmt"
	"net"
)

func (s *Server) Start() error {
	s.logger.Info("Server started on port " + s.port)

	port := fmt.Sprintf(":%s", s.port)

	listener, err := net.Listen("tcp", port)

	if err != nil {
		return domain.ErrServerStart
	}

	s.listener = listener

	err = s.cache.Start()

	if err != nil {
		return domain.ErrCacheStart
	}

	for {
		conn, err := s.listener.Accept()

		if err != nil {
			return domain.ErrServerAccept
		}

		go s.handle(conn)
	}
}

func (s *Server) Stop() error {
	err := s.listener.Close()

	if err != nil {
		return domain.ErrServerStop
	}

	s.logger.Info("Server stopped")

	return nil
}
