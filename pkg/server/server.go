package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port         int
	proxyHandler http.Handler
}

func New() Server {
	headers := &http.Header{}
	headers.Add("Proxied-From", "easyproxy")

	s := Server{
		port: 8080,
		proxyHandler: &Proxy{
			headers: *headers,
		},
	}

	return s
}

func (s *Server) Run() {
	// Specify the address and port to listen on
	addr := fmt.Sprintf(":%d", s.port)
	log.Printf("Starting server on %s\n", addr)
	// Start the HTTP server
	if err := http.ListenAndServe(addr, s.proxyHandler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
