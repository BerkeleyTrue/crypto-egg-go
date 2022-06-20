package config

import (
	"fmt"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
)

var Port string = "3000"
var GinReleaseMode string = "false"
var Version string = "dev"
var User string = "Anon"
var Time string = ""
var Hash string = ""

type (
	HTTP struct {
		Port           string
		GinReleaseMode bool
	}

	Config struct {
		HTTP    `yaml:"http"`
		Coins   []string `yaml:"coins"`
		Version string
		Hash    string
		Time    string
		User    string
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		return nil, fmt.Errorf("config err: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	cfg.HTTP.Port = Port
	cfg.Version = Version
  cfg.Hash = Hash
	cfg.Time = Time
	cfg.User = User

	ginReleaseMode, err := strconv.ParseBool(GinReleaseMode)

	if err != nil {
		return nil, err
	}

	cfg.HTTP.GinReleaseMode = ginReleaseMode

	return cfg, nil
}
