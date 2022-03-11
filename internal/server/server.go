package server

import (
	"context"
	"net/http"
	"time"
)

type Server interface {
	Run() error
	Shutdown(context.Context) error
}

const maxHeaderSize = 1 << 20 // 1Mb
var timeout time.Duration = 10 * time.Second

type server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler) *server {
	return &server{
		httpServer: &http.Server{
			Addr:           ":8080",
			Handler:        handler,
			MaxHeaderBytes: maxHeaderSize,
			ReadTimeout:    timeout,
			WriteTimeout:   timeout,
		},
	}
}

func (s *server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
