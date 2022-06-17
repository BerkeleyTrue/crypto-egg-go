package coingecko

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/berkeleytrue/crypto-egg-go/internal/core/domain"
	"gopkg.in/eapache/go-resiliency.v1/retrier"
	retry "gopkg.in/h2non/gentleman-retry.v2"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

var apiUrl = "https://api.coingecko.com/api/v3/"

type CGClient struct {
	client *gentleman.Client
}

type jsonData struct {
	URL     string            `json:"url"`
	Origin  string            `json:"origin"`
	Headers map[string]string `json:"headers"`
}

func Init() CGClient {
	client := gentleman.New()
	client.URL(apiUrl)
	return CGClient{
		client,
	}
}

func (c CGClient) Ping() (bool, error) {
	res, err := c.client.
		Path("/ping").
		Use(retry.New(retrier.New(retrier.ExponentialBackoff(100, 2000), nil))).
		Request().
		Send()

	if err != nil {
		return false, fmt.Errorf("Error making request: %w", err)
	}
	return res.Ok, nil
}

func (c CGClient) GetCoins(coins []string) ([]domain.Coin, error) {
	data, err := json.Marshal(coins)
	if err != nil {
		return nil, fmt.Errorf("Couldn't encode json: %w", err)
	}

	res, err := c.client.Path("/coins/markets").Method(http.MethodPost).Request().Use(body.JSON(data)).Send()
	if err != nil {
		return nil, fmt.Errorf("Couldn't complete request: %w", err)
	}
	if !res.Ok {
		return nil, fmt.Errorf("Invalid response from server: %w", err)
	}

	json := &jsonData{}
	res.JSON(&json)
	fmt.Printf("body: %#v\n", json)
	// TODO translate response to coins
	return nil, nil
}
