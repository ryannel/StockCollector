package models

import (
	"fmt"
	"stockCollector/2-core/outboundProviders"
)

func NewCorporation(symbol string, provider outboundProviders.StockProviderInterface) (Corporation, error) {
	companyInfo, err := provider.GetCompanyInfo(symbol)
	if err != nil {
		return Corporation{}, fmt.Errorf("error fetching company information for symbol: `%s` - Error: %w", symbol, err)
	}

	priceHistory, err := provider.GetStockPriceHistory(symbol)
	if err != nil {
		return Corporation{}, fmt.Errorf("error fetching stock price history for symbol: `%s` - Error: %w", symbol, err)
	}

	insiderActivity, err := provider.GetInsiderActivity(symbol)
	if err != nil {
		return Corporation{}, fmt.Errorf("error fetching insider activity for symbol: `%s` - Error: %w", symbol, err)
	}
	
	return Corporation{
		CompanyName: companyInfo.CompanyName,
		Industry:    companyInfo.Industry,
		Sector:      companyInfo.Sector,
		Stocks: []StockInfo{
			{
				Symbol:      symbol,
				Exchange:    companyInfo.Exchange,
				History:     priceHistory,
				InsiderActivity: insiderActivity,
			},
		},
	}, nil
}

type Corporation struct {
	CompanyName string
	Industry    string
	Sector      string
	Stocks      []StockInfo
}

type StockInfo struct {
	Symbol      string
	Exchange    string
	History     []outboundProviders.StockPriceSnapshot
	InsiderActivity []outboundProviders.InsiderActivity
}


