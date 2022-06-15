package app

import (
	"fmt"

	"github.com/berkeleytrue/crypto-agg-go/config"
)

func Run(cfg *config.Config) {
	fmt.Printf("Hello World %s", cfg)
}
