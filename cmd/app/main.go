package main

import (
	"log"

	"github.com/berkeleytrue/crypto-egg-go/config"
	"github.com/berkeleytrue/crypto-egg-go/internal/app"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	cfg, err := config.NewConfig()

	if err != nil {
		logger.Fatal(err)
	}

	log.Printf("Version: %s", cfg.Version)
	log.Printf("Hash: %s", cfg.Hash)
	log.Printf("Build User: %s", cfg.User)
	log.Printf("Build Time: %s", cfg.Time)

	app.Run(cfg)
}
