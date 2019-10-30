package alphaVantageProvider

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_AlphaVantageDailyRecent_GivenSymbol_ShouldNotError(t *testing.T) {
	sut, err := New("demo", http.Client{})
	assert.Nil(t, err)

	result, err := sut.GetDailyStockPriceHistory20d("MSFT")
	assert.Nil(t, err)

	assert.NotEqual(t, 0, len(result))

	assert.NotEqual(t, float64(0), result[0].Volume)
	assert.NotEqual(t, float64(0), result[0].High)
	assert.NotEqual(t, float64(0), result[0].Low)
	assert.NotEqual(t, float64(0), result[0].Open)
	assert.NotEqual(t, float64(0), result[0].Close)
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewMockHttpClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func Test_AlphaVantageDailyRecent_GivenSymbol_CallCorrectUrl(t *testing.T) {
	client := NewMockHttpClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY&symbol=MSFT", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	sut, err := New("demo", *client)
	assert.Nil(t, err)

	_, _ = sut.GetDailyStockPriceHistory20d("MSFT")
}

func Test_AlphaVantageDaily_GivenSymbol_CallCorrectUrl(t *testing.T) {
	client := NewMockHttpClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://www.alphavantage.co/query?apikey=demo&function=TIME_SERIES_DAILY&outputsize=full&symbol=MSFT", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	sut, err := New("demo", *client)
	assert.Nil(t, err)

	_, _ = sut.GetDailyStockPriceHistory20y("MSFT")
}