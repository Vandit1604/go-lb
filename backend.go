package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	mux              sync.RWMutex // Use RWMutex for read operations
	url              *url.URL
	alive            bool
	aliveConnections int
	reverseProxy     *httputil.ReverseProxy
}

func NewBackend(u *url.URL, rp *httputil.ReverseProxy) *Backend {
	return &Backend{
		url:          u,
		alive:        true,
		reverseProxy: rp,
	}
}

// IsAlive checks if the backend is alive.
func (b *Backend) IsAlive() bool {
	b.mux.RLock()
	defer b.mux.RUnlock()

	return b.alive
}

// SetAlive sets the alive status of the backend.
func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	defer b.mux.Unlock()

	b.alive = alive
}

// GetURL returns the URL of the backend.
func (b *Backend) GetURL() *url.URL {
	b.mux.RLock()
	defer b.mux.RUnlock()

	return b.url
}

// GetActiveConnections returns the number of active connections.
func (b *Backend) GetActiveConnections() int {
	b.mux.RLock()
	defer b.mux.RUnlock()

	return b.aliveConnections
}

// Serve handles HTTP requests by forwarding them to the backend server.
func (b *Backend) Serve(rw http.ResponseWriter, req *http.Request) {
	b.reverseProxy.ServeHTTP(rw, req)
	b.mux.Lock()
	b.aliveConnections++
	b.mux.Unlock()
}
