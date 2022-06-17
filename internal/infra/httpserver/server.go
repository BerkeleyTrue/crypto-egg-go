package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/berkeleytrue/crypto-agg-go/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
	notify chan error
}

func (s *Server) Start() func() {

	go func() {
		fmt.Printf("Starting Server on %s\n", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.notify <- err
		}
	}()

	return func() {
		fmt.Println("Shutdown requested")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		defer func() {
			cancel()
			close(s.notify)
		}()

		s.server.SetKeepAlivesEnabled(false)
		s.notify <- s.server.Shutdown(ctx)
		fmt.Println("Shutdown complete")
	}
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func New(handler *gin.Engine, cfg *config.HTTP) *Server {
	httpServer := &http.Server{
		Addr:    ":" + string(cfg.Port),
		Handler: handler,
	}

	s := &Server{
		server: httpServer,
		notify: make(chan error, 1),
	}

	return s
}
