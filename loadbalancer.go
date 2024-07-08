package main

import (
	"log"
	"net/http"
)

type LoadBalancer struct {
	serverPool ServerPool
}

func InitLoadBalancer(serverPool ServerPool) *LoadBalancer {
	return &LoadBalancer{
		serverPool: serverPool,
	}
}

// Serve handles incoming HTTP requests and forwards them to an available backend server.
func (lb *LoadBalancer) Serve(w http.ResponseWriter, r *http.Request) {
	backend := lb.serverPool.GetNextValidServer()

	if backend == nil {
		log.Println("No available servers")
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	log.Println("Forwarding request to backend:", backend.GetURL())
	backend.Serve(w, r)
}
