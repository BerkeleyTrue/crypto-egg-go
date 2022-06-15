package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/berkeleytrue/crypto-agg-go/config"
)

type Server struct {
	server *http.Server
	wg     *sync.WaitGroup
}

func Start(s *Server) {
	s.wg.Add(1)
	go func() {
		fmt.Printf("Starting Server on %s\n", s.server.Addr)
		err := s.server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func New(handler http.Handler, cfg *config.HTTP, wg *sync.WaitGroup) *Server {
	httpServer := &http.Server{
		Addr:    ":" + string(cfg.Port),
		Handler: handler,
	}

	s := &Server{
		server: httpServer,
		wg:     wg,
	}

	return s
}
