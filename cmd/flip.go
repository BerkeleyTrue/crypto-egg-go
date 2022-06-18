/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// flipCmd represents the flip command
var flipCmd = &cobra.Command{
	Use:   "flip",
	Short: "Get the current flippening ratio",
	Long:  `Get the current flippening ratio of Ethereum to Bitcoin by marketCap`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flip called")
	},
}

func init() {
	rootCmd.AddCommand(flipCmd)
}
