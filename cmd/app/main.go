package main

import (
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

	app.Run(cfg)
}
