package iexCloudProvider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"stockCollector/2-core/outboundProviders"
	"stockCollector/infrastructure/helpers"
)

var baseUrl = "https://cloud.iexapis.com/stable/"

func New(apiToken string, httpClient http.Client) (IexCloudProvider, error) {
	if apiToken == "" {
		return IexCloudProvider{}, errors.New("IEX Cloud API token must be provided")
	}

	requestBuilder, err := helpers.NewUrlBuilder(baseUrl)
	if err != nil {
		return IexCloudProvider{}, fmt.Errorf("unable to create url builder- %w", err)
	}
	requestBuilder.AddQueryParameter("token", apiToken)

	return IexCloudProvider{
		baseUrl: requestBuilder.GetUrl(),
		httpClient:     httpClient,
	}, nil
}

type IexCloudProvider struct {
	httpClient     http.Client
	baseUrl string
}

func (provider IexCloudProvider) GetCompanyInfo(symbol string) (outboundProviders.CompanyInfo, error) {
	urlBuilder, err := helpers.NewUrlBuilder(provider.baseUrl)
	if err != nil {
		return outboundProviders.CompanyInfo{}, fmt.Errorf("error creating URL builder for IexCloud provider - %s : %w", provider.baseUrl, err)
	}
	urlBuilder.AppendPath(fmt.Sprintf("stock/%s/company", symbol))

	requestUrl := urlBuilder.GetUrl()
	responseBody, err := provider.httpGet(requestUrl)
	if err != nil {
		return outboundProviders.CompanyInfo{}, err
	}

	var responseData CompanyDto
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return outboundProviders.CompanyInfo{}, fmt.Errorf("unable to parse json httpResponse from IEX Cloud - Error: %w", err)
	}

	return mapCompany(responseData)
}

func (provider IexCloudProvider) GetDailyStockPriceHistory1y(symbol string) ([]outboundProviders.StockPriceSnapshot, error) {
	return provider.getDailyStockPriceHistory(symbol, "1y")
}

func (provider IexCloudProvider) GetDailyStockPriceHistory20y(symbol string) ([]outboundProviders.StockPriceSnapshot, error) {
	return provider.getDailyStockPriceHistory(symbol, "20y")
}

func (provider IexCloudProvider) GetDailyStockPriceHistory20d(symbol string) ([]outboundProviders.StockPriceSnapshot, error) {
	return provider.getDailyStockPriceHistory(symbol, "20d")
}

func (provider IexCloudProvider) getDailyStockPriceHistory(symbol string, timeSpan string) ([]outboundProviders.StockPriceSnapshot, error) {
	urlBuilder, err := helpers.NewUrlBuilder(provider.baseUrl)
	if err != nil {
		return nil, fmt.Errorf("error creating URL builder for IexCloud provider - %s : %w", provider.baseUrl, err)
	}
	urlBuilder.AppendPath(fmt.Sprintf("stock/%s/chart/%s", symbol, timeSpan))

	requestUrl := urlBuilder.GetUrl()
	responseBody, err := provider.httpGet(requestUrl)
	if err != nil {
		return nil, err
	}

	var responseData []StockPriceSnapshotDto
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, fmt.Errorf("unable to parse json httpResponse from IEX Cloud - Error: %w", err)
	}

	return mapStockPriceSnapshots(responseData)
}

func (provider IexCloudProvider) httpGet(requestUrl string) ([]byte, error) {
	httpResponse, err := provider.httpClient.Get(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("IEX Cloud error sending request - %s: %w", requestUrl, err)
	}
	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("IEX Cloud service error %s: %s", requestUrl, httpResponse.Status)
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read IEX Cloud httpResponse body - Error : %w", err)
	}

	return responseBody, nil
}
