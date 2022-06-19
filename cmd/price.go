/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/utils/formatutil"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/h2non/gentleman.v2"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price [sym]",
	Short: "Get the price of a coin",
	Long:  `Get the current price of a coin`,
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		logger := zap.NewExample().Sugar()
		defer logger.Sync()

		if client := cmd.Context().Value("client"); client == nil {
			logger.Fatalf("Cmd doesn't have access to client: %#v", client)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		logger := zap.NewExample().Sugar()
		defer logger.Sync()

		client := cmd.Context().Value("client").(*gentleman.Client)
		sym := args[0]
		// fmt.Printf("price called w/ %s\n", sym)

		request := client.Request()

		request.AddPath("/api/coins/sym/:sym")
		request.Param("sym", sym)

		res, err := request.Send()

		if err != nil {
			logger.Fatalf("Request failed: %w", err)
		}

		if !res.Ok && res.StatusCode != 404 {
			logger.Fatalf("Bad response: %d", res.StatusCode)
		}
		if res.StatusCode == 404 {
			fmt.Println("0.00")
			return
		}

		coin := domain.Coin{}
		err = res.JSON(&coin)
		if err != nil {
			logger.Fatalf("Couldn't parse response: %w", err)
		}

		formatutil.PrintPrice(coin.Price)
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
}
