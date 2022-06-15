package app

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/berkeleytrue/crypto-agg-go/config"
	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

type Server struct {
	server *http.Server
}

func (s *Server) start() {
	wg.Add(1)
	go func() {
		s.server.ListenAndServe()
	}()
}

func Run(cfg *config.Config) {
	handler := gin.Default()
	handler.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	httpServer := &http.Server{
		Handler: handler,
	}

	httpServer.Addr = net.JoinHostPort("", cfg.HTTP.Port)

	s := &Server{
		server: httpServer,
	}

	s.start()
	wg.Wait()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

}
