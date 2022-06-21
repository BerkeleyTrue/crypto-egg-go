package config

import "time"

var Port string = "3000"
var Release string = "development"
var User string = "Anon"
var Time string = time.Now().Format(time.RFC3339);
var Hash string = "N/A"

type (
	HTTP struct {
		Port    string
	}

	Config struct {
		HTTP    `yaml:"http"`
		Coins   []string `yaml:"coins"`
		Hash    string
		Time    string
		User    string
		Release string
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	cfg.HTTP.Port = Port
	cfg.Hash = Hash
	cfg.Time = Time
	cfg.User = User
	cfg.Coins = []string{
		"ethereum",
		"bitcoin",
		"tezos",
		"pickle-finance",
		"olympus",
		"ethereum-name-service",
		"staked-ether",
	}

	cfg.Release = Release

	return cfg, nil
}
