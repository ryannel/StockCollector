package alphaVantageProvider

import (
	"math"
	"stockCollector/2-core/outboundProviders"
	"strconv"
)

func mapResponse(src TimeSeriesDailyDto) ([]outboundProviders.StockPriceSnapshot,  error) {
	result := make([]outboundProviders.StockPriceSnapshot, len(src.TimeSeriesDaily))

	i := 0
	for key, value := range src.TimeSeriesDaily {
		open, err :=  strconv.ParseFloat(value.Open, 64)
		if err != nil {
			return nil, err
		}

		high, err :=  strconv.ParseFloat(value.High, 64)
		if err != nil {
			return nil, err
		}

		low, err :=  strconv.ParseFloat(value.Low, 64)
		if err != nil {
			return nil, err
		}

		close, err :=  strconv.ParseFloat(value.Close, 64)
		if err != nil {
			return nil, err
		}

		volume, err :=  strconv.ParseFloat(value.Volume, 64)
		if err != nil {
			return nil, err
		}

		result[i].DateTime = key.Time
		result[i].Open = open
		result[i].High = high
		result[i].Low = low
		result[i].Close = close
		result[i].Volume = int(math.Round(volume))

		i++
	}

	return result, nil
}