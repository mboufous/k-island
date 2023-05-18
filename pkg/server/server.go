package server

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/mboufous/k-island/utils/log"
	"github.com/mboufous/k-island/utils/signal"
)

type Config struct {
	Host     string
	Port     string
	Handler  http.Handler
	Shutdown chan os.Signal
}

type Server struct {
	config *Config
	*http.Server
}

func New(config *Config) *Server {

	return &Server{
		config: config,
		Server: &http.Server{
			Handler: config.Handler,
			Addr:    net.JoinHostPort(config.Host, config.Port),
		},
	}
}

func (s *Server) listenAndServe() error {
	// Listen
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("could not listen on %s: %w", s.Addr, err)
	}
	// Server
	return s.Serve(listener)
}

func (s *Server) Start() {
	go func() {
		if err := s.listenAndServe(); err != http.ErrServerClosed {
			log.Error("server startup: %s", err)
			signal.Stop.Stop()
		}
	}()
	log.Infof("listening on http://%s/", s.Addr)
}
