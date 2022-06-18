/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price [sym]",
	Short: "Get the price of a coin",
	Long:  `Get the current price of a coin`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sym := args[0]
		fmt.Printf("price called w/ %s\n", sym)
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
}
