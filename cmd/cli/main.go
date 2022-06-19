package main

import (
	"github.com/berkeleytrue/crypto-egg-go/cmd"
	"github.com/berkeleytrue/crypto-egg-go/config"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	cfg, err := config.NewConfig()

	if err != nil {
		logger.Fatal(err)
	}
	cmd.Execute(cfg)
}
