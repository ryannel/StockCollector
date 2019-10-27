package alphaVantageProvider

import (
	"stockCollector/2-core/outboundProviders"
	"stockCollector/3-outbound/alphaVantageProvider/dto"
	"strconv"
)

func mapResponse(src dto.TimeSeriesDailyDto) (outboundProviders.StockPriceSeries,  error) {
	result := outboundProviders.StockPriceSeries{}

	for key, value := range src.TimeSeriesDaily {
		open, err :=  strconv.ParseFloat(value.Open, 64)
		if err != nil {
			return outboundProviders.StockPriceSeries{}, err
		}

		high, err :=  strconv.ParseFloat(value.High, 64)
		if err != nil {
			return outboundProviders.StockPriceSeries{}, err
		}

		low, err :=  strconv.ParseFloat(value.Low, 64)
		if err != nil {
			return outboundProviders.StockPriceSeries{}, err
		}

		close, err :=  strconv.ParseFloat(value.Close, 64)
		if err != nil {
			return outboundProviders.StockPriceSeries{}, err
		}

		volume, err :=  strconv.ParseFloat(value.Volume, 64)
		if err != nil {
			return outboundProviders.StockPriceSeries{}, err
		}

		result[key.Time] = outboundProviders.StockPriceSnapshot{
			Open:   open,
			High:   high,
			Low:    low,
			Close:  close,
			Volume: volume,
		}
	}

	return result, nil
}