package cmd

import (
	"fmt"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/berkeleytrue/crypto-egg-go/internal/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/h2non/gentleman.v2"
)

var stethCommand = &cobra.Command{
	Use:   "steth",
	Short: "Get the current steth discount",
	Long:  `Get the current current discount on steth`,
	Args:  cobra.MinimumNArgs(0),
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

		request := client.Request()

		request.AddPath("/api/coins")

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

		coins := []domain.Coin{}

		err = res.JSON(&coins)

		if err != nil {
			logger.Fatalf("Couldn't parse response: %w", err)
		}

    if len(coins) <= 0 {
      fmt.Println("0")
      return
    }

    var eth domain.Coin
    var steth domain.Coin

    for _, coin := range coins {
      if coin.Symbol == "eth" {
        eth = coin
      }

      if coin.Symbol == "steth" {
        steth = coin
      }
    }

    discount := (eth.Price - steth.Price) / eth.Price * 100

		utils.Formatter.PrintPrice(discount)
	},
}

func init() {
	rootCmd.AddCommand(stethCommand)
}
