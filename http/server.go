package http

import (
	"log"
	"net"
	"net/http"
)

// DefaultAddr is the default bind address.
const DefaultAddr = "localhost:3333"

// Server represents an HTTP server.
type Server struct {
	ln net.Listener

	// Handler to serve.
	Handler *Handler

	// Bind address to open.
	Addr string
}

// NewServer returns a new instance of Server.
func NewServer(handler *Handler) *Server {
	return &Server{
		Addr:    DefaultAddr,
		Handler: handler,
	}
}

// ListenAndServe listens and serves on address
func (s *Server) ListenAndServe() {
	log.Fatal(http.ListenAndServe(s.Addr, s.Handler))
}
