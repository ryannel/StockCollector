package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"stockCollector/2-core/models"
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

	log.Println("Fetching Nasdaq symbols")
	nasdaq := nasdaqProvider.New(http.Client{})
	symbols, err := nasdaq.GetAllSymbols()
	if err != nil {
		log.Panic(err)
	}
	log.Println("loaded " + strconv.Itoa(len(symbols)) + " symbols")

	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}


	outputFilePath := filepath.Join(wd, "nasdaqTradedCompanies.jsonl")
	log.Println("Generating output file: " + outputFilePath)
	jsonWriter, err := writer.NewJsonLinesWriter(outputFilePath)
	if err != nil {
		log.Panic(err)
	}
	defer jsonWriter.Close()

	for i, symbol := range symbols {
		progress := fmt.Sprintf("%f", float64(i+1)/float64(len(symbols)) * 100)
		log.Println("processing symbol number " + strconv.Itoa(i + 1) + " (" + symbol + ") - " + progress + "% complete")

		corp, err := models.NewCorporation(symbol, &nasdaq)
		if err != nil {
			log.Println(err)
		}

		err = jsonWriter.AppendLine(corp)
		if err != nil {
			log.Panic(err)
		}

		r := rand.Intn(3)
		time.Sleep(time.Duration(r) * time.Second)
	}
}

type Result struct {
	Corporation models.Corporation
	Err error
}

func paralelFetch(symbols []string, concurrency int) ([]models.Corporation, error) {
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

	companies := make([]models.Corporation, numSymbols)
	for i := 1; i <= numSymbols; i++ {
		result := <- resultsChan
		if result.Err != nil {
			log.Println("error loading price history for: " + result.Corporation.Stocks[0].Symbol + " - " + result.Err.Error())
		} else {
			companies = append(companies, result.Corporation)
		}
	}

	return companies, nil
}

func worker(jobChan <- chan string, resultChan chan <- Result) {
	nasdaq := nasdaqProvider.New(http.Client{})

	for symbol := range jobChan {
		log.Println("fetching price history for: " + symbol)
		corp, err := models.NewCorporation(symbol, &nasdaq)

		resultChan <- Result{
			Corporation: corp,
			Err: err,
		}

		r := rand.Intn(3)
		log.Println("Sleeping for: " + strconv.Itoa(r) + " seconds")
		time.Sleep(time.Duration(r) * time.Second)
	}
}