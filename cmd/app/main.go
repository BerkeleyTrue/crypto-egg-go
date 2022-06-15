package main

import (
	"log"
	"sync"

	"github.com/berkeleytrue/crypto-agg-go/config"
	"github.com/berkeleytrue/crypto-agg-go/internal/app"
)

var wg sync.WaitGroup

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg, &wg)
}
