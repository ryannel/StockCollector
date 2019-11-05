package config

import (
	"github.com/caarlos0/env"
)

type AppConfig struct {
	AlphaVantage *AlphaVantageConfig
	IexCloud     *IexCloud
}

type AlphaVantageConfig struct {
	ApiKey string `env:"ALPHA_VANTAGE_API_KEY,required"`
}

type IexCloud struct {
	ApiKey string `env:"IEXCLOUD_API_KEY,required"`
}

func ParseConfig() (*AppConfig, error) {
	conf := &AppConfig{
		AlphaVantage: &AlphaVantageConfig{},
		IexCloud:     &IexCloud{},
	}
	err := env.Parse(conf)

	return conf, err
}
