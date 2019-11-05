package main

import (
	"fmt"
	"log"
	"net/http"
	"stockCollector/2-core/outboundProviders"
	"stockCollector/3-outbound/nasdaqProvider"
)

func main() {
	//conf, err := config.ParseConfig()
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//stockCollector, err := alphaVantageProvider.New(conf.AlphaVantage.ApiKey, http.Client{})
	//if err != nil {
	//	log.Panic(err)
	//}

	nasdaq := nasdaqProvider.New(http.Client{})

	symbols, err := nasdaq.GetAllSymbols()
	if err != nil {
		log.Panic(err)
	}

	var companies []outboundProviders.Company
	for _, symbol := range symbols {
		fmt.Println("Fetching price history for: " + symbol)
		history, err := nasdaq.GetPriceHistory(symbol)
		if err != nil {
			log.Panic(err)
		}

		companies = append(companies, outboundProviders.Company{
			Symbol:       symbol,
			PriceHistory: history,
		})
	}

	//err = inboundServices.CollectStock(&stockCollector)
	//if err != nil {
	//	log.Panic(err)
	//}

	fmt.Println("Hello World")
}