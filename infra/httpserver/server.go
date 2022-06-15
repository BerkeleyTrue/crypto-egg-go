package httpserver

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/berkeleytrue/crypto-agg-go/config"
)

type Server struct {
	server *http.Server
	wg     *sync.WaitGroup
	notify chan error
}

func (s *Server) Start() {
	s.wg.Add(1)
	go func() {
		fmt.Printf("Starting Server on %s\n", s.server.Addr)
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func New(handler http.Handler, cfg *config.HTTP, wg *sync.WaitGroup) *Server {
	httpServer := &http.Server{
    Addr:    ":" + string(cfg.Port),
		Handler: handler,
	}

	s := &Server{
		server: httpServer,
		wg:     wg,
		notify: make(chan error, 1),
	}

	return s
}
