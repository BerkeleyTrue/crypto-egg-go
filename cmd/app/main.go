package main

import (
	"log"

	"github.com/berkeleytrue/crypto-agg-go/config"
	"github.com/berkeleytrue/crypto-agg-go/internal/app"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
