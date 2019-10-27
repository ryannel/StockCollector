package config

import (
	"github.com/caarlos0/env"
)

type AppConfig struct {
	AlphaVantage        *AlphaVantageConfig
}

type AlphaVantageConfig struct {
	ApiKey string `env:"ALPHA_VANTAGE_API_KEY,required"`
}

func ParseConfig() (*AppConfig, error) {
	conf := &AppConfig{
		AlphaVantage:  &AlphaVantageConfig{},
	}
	err := env.Parse(conf)

	return conf, err
}