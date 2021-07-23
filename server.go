package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	fmt.Printf("Run server on port : %s", port)
	s.server = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1mb

	}
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	fmt.Println("Server shutdown.")
	return s.server.Shutdown(ctx)
}
