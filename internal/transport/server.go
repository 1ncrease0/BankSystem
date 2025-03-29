package transport

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(address string, timeout time.Duration, idleTimeout time.Duration, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         address,
		Handler:      handler,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  idleTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
