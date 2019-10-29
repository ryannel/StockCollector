package iexCloudProvider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"stockCollector/2-core/outboundProviders"
	"stockCollector/3-outbound/iexCloudProvider/dto"
	"stockCollector/infrastructure/helpers"
)

var baseUrl = "https://cloud.iexapis.com/stable/"

func New(apiToken string, httpClient http.Client) (IexCloudProvider, error) {
	if apiToken == "" {
		return IexCloudProvider{}, errors.New("IEX Cloud API token must be provided")
	}

	requestBuilder, err := helpers.NewHttpRequestBuilder(baseUrl)
	if err != nil {
		return IexCloudProvider{}, fmt.Errorf("unable to create http request builder- %w", err)
	}
	requestBuilder.AddQueryParameter("token", apiToken)

	return IexCloudProvider{
		requestBuilder: requestBuilder,
		httpClient:     httpClient,
	}, nil
}

type IexCloudProvider struct {
	apiKey         string
	httpClient     http.Client
	requestBuilder helpers.HttpRequestBuilder
}

func (provider IexCloudProvider) GetDailyStockPriceHistory1y(symbol string) ([]outboundProviders.StockPriceSnapshot, error) {
	provider.requestBuilder.AppendPath(fmt.Sprintf("stock/%s/chart/1y", symbol))

	requestUrl := provider.requestBuilder.GetUrl()
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

	var responseData []dto.StockPriceSnapshotDto
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, fmt.Errorf("unable to parse json httpResponse from alpha vantage - Error: %w", err)
	}

	return mapStockPriceSnapshots(responseData)
}