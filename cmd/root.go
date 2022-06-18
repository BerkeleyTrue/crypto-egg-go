/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var CmdName string = "crypto-egg-go"

var rootCmd = &cobra.Command{
	Use:   CmdName,
	Short: "Get crypto prices and aggregates",
	Long: `Get crypto prices, gas price, and flippenings from various sources. For example:
  ` + CmdName + ` price [sym] => returns the current price of [sym]
  ` + CmdName + ` gas [sym] => returns the current price of [sym]
  ` + CmdName + ` flip => returns the current marketcap of Ethereum against Bitcoin
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
