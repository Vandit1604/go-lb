package main

import (
	"context"
	"net"
	"net/url"
	"sync"
)

type ServerPool struct {
	mux      sync.Mutex
	backends []*Backend // Change to slice of pointers for modification safety
	current  int
}

func (s *ServerPool) GetBackends() []*Backend {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.backends
}

// Adds backend in Serverpool
func (s *ServerPool) AddBackends(backend *Backend) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.backends = append(s.backends, backend)
}

// returns the next backend safely
func (s *ServerPool) getNextBackend() *Backend {
	s.mux.Lock()
	defer s.mux.Unlock()

	if len(s.backends) == 0 {
		return nil
	}

	s.current = (s.current + 1) % len(s.backends)
	return s.backends[s.current]
}

// GetNextValidServer finds and returns the next valid (alive) server.
func (s *ServerPool) GetNextValidServer() *Backend {
	for i := 0; i < len(s.backends); i++ {
		nextPeer := s.getNextBackend()
		if nextPeer != nil && nextPeer.IsAlive() {
			return nextPeer
		}
	}
	return nil
}

// IsBackendAlive checks if a backend server is alive by attempting a TCP connection.
func IsBackendAlive(ctx context.Context, url *url.URL) bool {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", url.Host)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}
