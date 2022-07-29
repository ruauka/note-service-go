package server

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, router *httprouter.Router) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        router,
			MaxHeaderBytes: 1 << 20, // 1 MB
			WriteTimeout:   time.Second * 10,
			ReadTimeout:    time.Second * 10,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
