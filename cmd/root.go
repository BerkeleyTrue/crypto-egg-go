/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"os"

	"github.com/berkeleytrue/crypto-egg-go/config"
	"github.com/spf13/cobra"
	"gopkg.in/h2non/gentleman.v2"
)

var CmdName string = "crypto-egg-go"
var rootCmd = &cobra.Command{
	Use:   CmdName,
	Short: "Get crypto prices and aggregates",
}

func Execute(cfg *config.Config) {

	rootCmd.Long = `Get crypto prices, gas price, and flippenings from various sources. For example:
  ` + CmdName + ` price [sym] => returns the current price of [sym]
  ` + CmdName + ` gas [sym] => returns the current price of [sym]
  ` + CmdName + ` flip => returns the current marketcap of Ethereum against Bitcoin

    Version: ` + cfg.Version + `
    Hash: ` + cfg.Hash + `
    Build Time: ` + cfg.Time + `
    Build User: ` + cfg.User + `
  `
	client := gentleman.New()
	client.URL("http://localhost:" + cfg.HTTP.Port)

	ctx := context.WithValue(context.Background(), "client", client)

	err := rootCmd.ExecuteContext(ctx)

	if err != nil {
		os.Exit(1)
	}
}
