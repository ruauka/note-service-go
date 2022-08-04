package server

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"web/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, router *httprouter.Router) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.App.Port,
			Handler:        router,
			MaxHeaderBytes: 1 << 20, // 1 MB
			WriteTimeout:   time.Second * time.Duration(cfg.WriteTimeout),
			ReadTimeout:    time.Second * time.Duration(cfg.ReadTimeout),
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
