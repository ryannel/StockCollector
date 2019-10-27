package alphaVantageProvider

import (
	"github.com/stretchr/testify/assert"
	"stockCollector/3-outbound/alphaVantageProvider/dto"
	"testing"
	"time"
)

func Test_MapResponse(t *testing.T) {
	key1 := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	key2 := time.Date(2020, 02, 39, 20, 34, 58, 651387237, time.UTC)
	timeSeries := make(map[dto.Iso8601]dto.StockPriceSnapshotDto)

	timeSeries[dto.Iso8601{Time: key1}] = dto.StockPriceSnapshotDto{
		Open:   "20",
		High:   "30",
		Low:    "12",
		Close:  "22",
		Volume: "400",
	}

	timeSeries[dto.Iso8601{Time: key2}] = dto.StockPriceSnapshotDto{
		Open:   "50.255121",
		High:   "400.5559",
		Low:    "2.2511",
		Close:  "60.2511",
		Volume: "1000.22",
	}

	src := dto.TimeSeriesDailyDto{
		MetaData:        dto.MetaDataDto{
			Information:   "Information",
			Symbol:        "Symbol",
			LastRefreshed: "LastRefreshed",
			OutputSize:    "OutputSize",
			TimeZone:      "TimeZone",
		},
		TimeSeriesDaily: timeSeries,
		Error:           "TestError",
	}

	response, err := mapResponse(src)
	assert.Nil(t, err)

	assert.Equal(t, float64(20), response[key1].Open)
	assert.Equal(t, float64(30), response[key1].High)
	assert.Equal(t, float64(12), response[key1].Low)
	assert.Equal(t, float64(22), response[key1].Close)
	assert.Equal(t, float64(400), response[key1].Volume)

	assert.Equal(t, 50.255121, response[key2].Open)
	assert.Equal(t, 400.5559, response[key2].High)
	assert.Equal(t, 2.2511, response[key2].Low)
	assert.Equal(t, 60.2511, response[key2].Close)
	assert.Equal(t, 1000.22, response[key2].Volume)
}