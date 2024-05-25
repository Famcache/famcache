package server

import (
	"famcache/domain/cache"
	"net"
)

// handleSet handles the SET command
// set <key> <value> <ttl>
func (s *Server) handleSet(conn net.Conn, parts []string) {
	if len(parts) != 3 {
		s.logger.Error("Invalid SET request")

		s.replyError(conn, "Invalid SET request")
	}

	key := parts[1]
	value := parts[2]

	var ttl *string = nil

	if len(parts) == 4 {
		ttl = &parts[3]
	}

	err := s.cache.Set(cache.SetOptions{
		Key:   key,
		Value: value,
		TTL:   ttl,
	})

	if err != nil {
		s.logger.Error("Error setting key")

		s.replyError(conn, err.Error())
	}

	s.logger.Info("SET", key, value)

	s.replySuccess(conn)
}
