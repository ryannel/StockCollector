package main

import (
	"fmt"
	"log"
	"net/http"
	"stockCollector/2-core/inboundServices"
	"stockCollector/3-outbound/alphaVantageProvider"
	"stockCollector/infrastructure/config"
)

func main() {
	conf, err := config.ParseConfig()
	if err != nil {
		log.Panic(err)
	}

	stockCollector, err := alphaVantageProvider.New(conf.AlphaVantage.ApiKey, http.Client{})
	if err != nil {
		log.Panic(err)
	}

	err = inboundServices.CollectStock(&stockCollector)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Hello World")
}