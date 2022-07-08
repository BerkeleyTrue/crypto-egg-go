package cmd

import (
	"fmt"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/h2non/gentleman.v2"
)

var gasCmd = &cobra.Command{
	Use:   "gas",
	Short: "Get current gas base fee for Ethereum",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logger := zap.NewExample().Sugar()
		defer logger.Sync()

		client := cmd.Context().Value("client").(*gentleman.Client)
		request := client.Request()

		request.AddPath("/api/gas")

		res, err := request.Send()

		if err != nil {
			logger.Fatal("Request failed: %w", err)
		}

		if !res.Ok && res.StatusCode != 404 {
			logger.Fatalf("Bad response: %d", res.StatusCode)
		}

		if res.StatusCode == 404 {
			fmt.Println("0")
			return
		}

		gas := domain.Gas{}
		err = res.JSON(&gas)
		if err != nil {
			logger.Fatal("Couldn't parse response: %w", err)
		}

		fmt.Printf("%.2f\n", gas.Base)

	},
}

func init() {
	rootCmd.AddCommand(gasCmd)
}
