package server

import (
	"net/http"
	"time"

	"github.com/Jamshid90/flight/internal/config"
)

type Server struct {
	config  *config.Config
	handler http.Handler
}

func NewServer(config *config.Config, handler http.Handler) *Server {
	return &Server{config, handler}
}

func (s *Server) Run() error {
	server := http.Server{
		Addr:         s.config.Server.Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      s.handler,
	}
	return server.ListenAndServe()
}
