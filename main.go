package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

func main() {
	config, err := LBConfig()
	if err != nil {
		log.Fatalf("Error getting the config: %v", err)
	}

	serverPool := ServerPool{
		mux:      sync.Mutex{},
		backends: []*Backend{}, // Use pointer to Backend for safer concurrent access
		current:  0,
	}

	for _, backendURL := range config.Backends {
		if url, err := url.Parse(backendURL); err != nil {
			log.Fatalf("Error parsing URL %s: %v", backendURL, err)
		} else {
			rp := httputil.NewSingleHostReverseProxy(url)
			backendServer := NewBackend(url, rp)
			serverPool.AddBackends(backendServer) // Pass pointer directly
		}
	}

	lb := InitLoadBalancer(serverPool)

	mux := http.NewServeMux()
	mux.HandleFunc("/", lb.Serve) // Directly assign the method without wrapping it

	log.Println("Starting load balancer on port :9000")
	if err := http.ListenAndServe(":9000", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
