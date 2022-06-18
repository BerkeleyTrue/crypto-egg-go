/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/berkeleytrue/crypto-egg-go/config"
	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/utils/formatutil"
	"github.com/spf13/cobra"
	"gopkg.in/h2non/gentleman.v2"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price [sym]",
	Short: "Get the price of a coin",
	Long:  `Get the current price of a coin`,
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if cfg := cmd.Context().Value("cfg"); cfg == nil {
			log.Fatalf("Cmd doesn't have access to ctx: %#v", cfg)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfg := cmd.Context().Value("cfg").(*config.Config)
		sym := args[0]
		// fmt.Printf("price called w/ %s\n", sym)

		client := gentleman.New()
		client.URL("http://localhost:" + cfg.HTTP.Port)

		request := client.Request()

		request.AddPath("/api/coins/sym/:sym")
		request.Param("sym", sym)

		res, err := request.Send()

		if err != nil {
			log.Fatal("Request failed: %w", err)
		}

		if !res.Ok && res.StatusCode != 404 {
			log.Fatalf("Bad response: %d", res.StatusCode)
		}
		if res.StatusCode == 404 {
			fmt.Println("0.00")
			return
		}

		coin := domain.Coin{}
		err = res.JSON(&coin)
		if err != nil {
			log.Fatal("Couldn't parse response: %w", err)
		}

		formatutil.PrintPrice(coin.Price)
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)
}
