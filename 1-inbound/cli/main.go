package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"stockCollector/2-core/outboundProviders"
	"stockCollector/3-outbound/nasdaqProvider"
	"stockCollector/3-outbound/writer"
	"strconv"
	"time"
)

type ComplexError struct {
	Message string
	Code    int
}

func (ce ComplexError) Format(f fmt.State, c rune) {
	_, _ = f.Write([]byte("test format"))
}

func (ce ComplexError) Error() string {
	return fmt.Sprint(ce)
}

func main() {
	//conf, err := config.ParseConfig()
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//stockCollector, err := alphaVantageProvider.NewJsonWriter(conf.AlphaVantage.ApiKey, http.Client{})
	//if err != nil {
	//	log.Panic(err)
	//}

	nasdaq := nasdaqProvider.New(http.Client{})
	symbols, err := nasdaq.GetAllSymbols()
	if err != nil {
		log.Panic(err)
	}

	companies, err := paralelFetch(symbols[0:5000], 1)

	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	outputWriter, err := writer.NewJsonWriter(wd)
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
		if result.Err != nil {
			log.Println("error loading price history for: " + result.Company.Symbol + " - " + result.Err.Error())
		} else {
			companies = append(companies, result.Company)
		}
	}

	return companies, nil
}

func worker(jobChan <- chan string, resultChan chan <- Result) {
	nasdaq := nasdaqProvider.New(http.Client{})

	for symbol := range jobChan {
		log.Println("fetching price history for: " + symbol)
		history, err := nasdaq.GetPriceHistory(symbol)

		company := outboundProviders.Company{
			Symbol:       symbol,
			PriceHistory: history,
		}

		resultChan <- Result{
			Company: company,
			Err: err,
		}

		r := rand.Intn(3)
		log.Println("Sleeping for: " + strconv.Itoa(r) + " seconds")
		time.Sleep(time.Duration(r) * time.Second)
	}
}