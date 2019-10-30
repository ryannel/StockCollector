package alphaVantageProvider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"stockCollector/2-core/outboundProviders"
	"stockCollector/infrastructure/helpers"
)

func New(apiKey string, httpClient http.Client) (AlphaVantageProvider, error) {
	if apiKey == "" {
		return AlphaVantageProvider{}, errors.New("alpha vantage API key must be provided")
	}

	requestBuilder, err := helpers.NewUrlBuilder("https://www.alphavantage.co/query")
	if err != nil {
		return AlphaVantageProvider{}, fmt.Errorf("unable to create http request builder- %w", err)
	}
	requestBuilder.AddQueryParameter("apikey", apiKey)

	return AlphaVantageProvider{
		baseUrl:    requestBuilder.GetUrl(),
		httpClient: httpClient,
	}, nil
}

type AlphaVantageProvider struct {
	httpClient http.Client
	baseUrl    string
}

func (provider *AlphaVantageProvider) GetDailyStockPriceHistory20y(symbol string) ([]outboundProviders.StockPriceSnapshot, error) {
	requestBuilder, err := helpers.NewUrlBuilder(provider.baseUrl)
	if err != nil {
		return nil, err
	}
	requestBuilder.AddQueryParameter("outputsize", "full")
	return provider.getTimeSeriesDaily(requestBuilder, symbol)
}

func (provider *AlphaVantageProvider) GetDailyStockPriceHistory20d(symbol string) ([]outboundProviders.StockPriceSnapshot, error) {
	requestBuilder, err := helpers.NewUrlBuilder(provider.baseUrl)
	if err != nil {
		return nil, err
	}
	return provider.getTimeSeriesDaily(requestBuilder, symbol)
}

func (provider *AlphaVantageProvider) getTimeSeriesDaily(requestBuilder helpers.UrlBuilder, symbol string) ([]outboundProviders.StockPriceSnapshot, error) {
	requestBuilder.AddQueryParameter("symbol", symbol)
	requestBuilder.AddQueryParameter("function", "TIME_SERIES_DAILY")

	requestUrl := requestBuilder.GetUrl()
	httpResponse, err := provider.httpClient.Get(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("alpha vantage error sending request - %s: %w", requestUrl, err)
	}
	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("alpha vantage service error %s: %s", requestUrl, httpResponse.Status)
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read alpha vantage httpResponse body - Error : %w", err)
	}

	responseData := TimeSeriesDailyDto{}
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, fmt.Errorf("unable to parse json httpResponse from alpha vantage - Error: %w", err)
	}

	if responseData.Error != "" {
		return nil, fmt.Errorf("error in Alpha Vantage httpResponse - Error: %s", responseData.Error)
	}

	// Information only seems to be populated when results can't be generated, treating this as an error for now
	if responseData.Information != "" {
		return nil, fmt.Errorf("error in Alpha Vantage httpResponse - Error: %s", responseData.Information)
	}

	mappedResponse, err := mapResponse(responseData)
	if err != nil {
		return nil, fmt.Errorf("mapping alpha vantage httpResponse - Error: %s", responseData.Information)
	}

	return mappedResponse, nil
}
