package iexCloudProvider

import (
	"stockCollector/2-core/outboundProviders"
	"time"
)

func mapStockPriceSnapshots(src []StockPriceSnapshotDto) ([]outboundProviders.StockPriceSnapshot,  error) {
	results := make([]outboundProviders.StockPriceSnapshot, len(src))

	for i, value := range src {
		date, err := time.Parse("2006-01-02", value.Date)
		if err != nil {
			return nil, err
		}

		results[i] = outboundProviders.StockPriceSnapshot{
			DateTime: date,
			Open:     value.Open,
			High:     value.High,
			Low:      value.Low,
			Close:    value.Close,
			Volume:   value.Volume,
		}
	}

	return results, nil
}

func mapCompany(company CompanyDto) (outboundProviders.CompanyInfo, error) {
	return outboundProviders.CompanyInfo{
		CompanyName:  company.CompanyName,
		Industry:     company.Industry,
		Sector:       company.Sector,
		Symbol:       company.Symbol,
		Exchange:     company.Exchange,
		Cusip:        "",
		PriceHistory: nil,
	}, nil
}