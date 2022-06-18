package coingecko

import (
	"fmt"
	"strings"
	"time"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"github.com/jinzhu/copier"
	"gopkg.in/eapache/go-resiliency.v1/retrier"
	retry "gopkg.in/h2non/gentleman-retry.v2"
	"gopkg.in/h2non/gentleman.v2"
)

var apiUrl = "https://api.coingecko.com/api/v3/"

type CGClient struct {
	client *gentleman.Client
}

type jsonData struct {
	ID        string `json:"id"`
	Symbol    string
	Name      string
	Price     float32 `json:"current_price"`
	MarketCap int64   `json:"market_cap"`
}

type jsonRes []jsonData

func Init() CGClient {
	client := gentleman.New()
	client.URL(apiUrl)
	return CGClient{
		client,
	}
}

func (c CGClient) Ping() (bool, error) {
	res, err := c.client.
		Request().
		AddPath("ping").
		Use(retry.New(retrier.New(retrier.ExponentialBackoff(100, 2000), nil))).
		Send()

	if err != nil {
		return false, fmt.Errorf("Error making request: %w", err)
	}
	if !res.Ok {
		fmt.Printf("ping not ok error: %d\n", res.StatusCode)
	}
	return res.Ok, nil
}

func (c CGClient) GetCoins(ids []string) ([]domain.Coin, error) {
	request := c.client.
		Request().
		AddPath("coins/markets").
		SetQuery("vs_currency", "usd").
		SetQuery("ids", strings.Join(ids, ", "))

	res, err := request.Send()
	if err != nil {
		return nil, fmt.Errorf("Couldn't complete request: %w", err)
	}
	if !res.Ok {
		return nil, fmt.Errorf("Invalid response from server: %d", res.StatusCode)
	}

	json := jsonRes{}
	err = res.JSON(&json)
	if err != nil {
		return nil, fmt.Errorf("Couldn't decode response: %w", err)
	}

  numOfRes := len(json)
	coins := make([]domain.Coin, numOfRes)

	if numOfRes >= 1 {
		for idx, item := range json {
			coin := domain.Coin{UpdatedAt: time.Now()}
			copier.Copy(&coin, &item)
			coins[idx] = coin
		}
	}

	return coins, nil
}
