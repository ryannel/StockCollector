package alphaVantageProvider

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_MapResponse(t *testing.T) {
	key1 := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	key2 := time.Date(2020, 02, 39, 20, 34, 58, 651387237, time.UTC)
	timeSeries := make(map[Iso8601]StockPriceSnapshotDto)

	timeSeries[Iso8601{Time: key1}] = StockPriceSnapshotDto{
		Open:   "20",
		High:   "30",
		Low:    "12",
		Close:  "22",
		Volume: "400",
	}

	timeSeries[Iso8601{Time: key2}] = StockPriceSnapshotDto{
		Open:   "50.255121",
		High:   "400.5559",
		Low:    "2.2511",
		Close:  "60.2511",
		Volume: "1000.22",
	}

	src := TimeSeriesDailyDto{
		MetaData:        MetaDataDto{
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

	for _, value := range response {
		if value.DateTime == key1 {
			assert.Equal(t, float64(20), response[0].Open)
			assert.Equal(t, float64(30), response[0].High)
			assert.Equal(t, float64(12), response[0].Low)
			assert.Equal(t, float64(22), response[0].Close)
			assert.Equal(t, 400, response[0].Volume)
		} else {
			assert.Equal(t, 50.255121, response[1].Open)
			assert.Equal(t, 400.5559, response[1].High)
			assert.Equal(t, 2.2511, response[1].Low)
			assert.Equal(t, 60.2511, response[1].Close)
			assert.Equal(t, 1000, response[1].Volume)
		}
	}
}