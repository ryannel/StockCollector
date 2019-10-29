package iexCloudProvider

import (
	"stockCollector/2-core/outboundProviders"
	"stockCollector/3-outbound/iexCloudProvider/dto"
	"time"
)

func mapStockPriceSnapshots(src []dto.StockPriceSnapshotDto) ([]outboundProviders.StockPriceSnapshot,  error) {
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