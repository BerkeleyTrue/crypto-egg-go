/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gasCmd represents the gas command
var gasCmd = &cobra.Command{
	Use:   "gas",
	Short: "Get current gas base fee for Ethereum",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gas called")
	},
}

func init() {
	rootCmd.AddCommand(gasCmd)
}
