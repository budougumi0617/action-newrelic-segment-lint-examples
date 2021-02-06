package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// Server is graceful start server.
type Server struct {
	Listener net.Listener
	Server   *http.Server
}

// Start to accept listener.
func (s *Server) Start() error {
	return s.Server.Serve(s.Listener)
}

// Shutdown terminates server by graceful shutdown.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

func SampleHandler(w http.ResponseWriter, req *http.Request) {
	defer newrelic.FromContext(req.Context()).StartSegment("sample_handler").End()
	fmt.Fprintf(w, "Hello, %q", req.URL.Path)
}

func SampleHandler2(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %q", req.URL.Path)
}
