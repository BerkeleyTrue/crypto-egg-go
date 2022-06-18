package main

import (
	"log"

	"github.com/berkeleytrue/crypto-egg-go/cmd"
	"github.com/berkeleytrue/crypto-egg-go/config"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute(cfg)
}
