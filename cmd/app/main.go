package main

import (
	"log"

	"github.com/berkeleytrue/crypto-egg-go/config"
	"github.com/berkeleytrue/crypto-egg-go/internal/app"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
