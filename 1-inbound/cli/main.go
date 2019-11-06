package main

import (
	"log"
	"net/http"
	"os"
	"stockCollector/2-core/outboundProviders"
	"stockCollector/3-outbound/nasdaqProvider"
	writer "stockCollector/3-outbound/writer"
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

	companies, err := paralelFetch(symbols[0:5000], 5)

	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	outputWriter, err := writer.New(wd)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Printing output file")
	err = outputWriter.WriteJson(companies, "nasdaqTradedCompanies.json")
	if err != nil {
		log.Panic(err)
	}
}

type Result struct {
	Company outboundProviders.Company
	Err error
}

func paralelFetch(symbols []string, concurrency int) ([]outboundProviders.Company, error) {
	numSymbols := len(symbols)
	jobsChan := make(chan string, numSymbols)
	resultsChan := make(chan Result, numSymbols)

	for workerId := 1; workerId <= concurrency; workerId++ {
		go worker(jobsChan, resultsChan)
	}

	for _, symbol := range symbols {
		jobsChan <- symbol
	}
	close(jobsChan)

	companies := make([]outboundProviders.Company, numSymbols)

	for i := 1; i <= numSymbols; i++ {
		result := <- resultsChan
		if result.Err == nil {
			companies = append(companies, result.Company)
			//return nil, result.Err
		}
	}

	return companies, nil
}

func worker(jobChan <- chan string, results chan <- Result) {
	nasdaq := nasdaqProvider.New(http.Client{})

	for symbol := range jobChan {
		log.Println("fetching price history for: " + symbol)
		history, err := nasdaq.GetPriceHistory(symbol)

		if err != nil {
			log.Println("error loading price history for: " + symbol + " - " + err.Error())
		}

		company := outboundProviders.Company{
			Symbol:       symbol,
			PriceHistory: history,
		}

		results <- Result{
			Company: company,
			Err: err,
		}
	}
}