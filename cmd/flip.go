/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/spf13/cobra"
	"gopkg.in/h2non/gentleman.v2"
)

// flipCmd represents the flip command
var flipCmd = &cobra.Command{
	Use:   "flip",
	Short: "Get the current flippening ratio",
	Long:  `Get the current flippening ratio of Ethereum to Bitcoin by marketCap`,
	Run: func(cmd *cobra.Command, args []string) {
		client := cmd.Context().Value("client").(*gentleman.Client)

		request := client.Request()

		request.AddPath("/api/flippening")

		res, err := request.Send()

		if err != nil {
			log.Fatal("Request failed: %w", err)
		}

		if !res.Ok && res.StatusCode != 404 {
			log.Fatalf("Bad response: %d", res.StatusCode)
		}
		if res.StatusCode == 404 {
			fmt.Println("0")
			return
		}

		flip := domain.Flippening{}
		err = res.JSON(&flip)
		if err != nil {
			log.Fatal("Couldn't parse response: %w", err)
		}

		fmt.Printf("%.2f\n", flip.Ratio)
	},
}

func init() {
	rootCmd.AddCommand(flipCmd)
}
