package nasdaqProvider

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNasdaqProvider_GetAllSymbols_ShouldNotBeEmpty(t *testing.T) {
	sut := New(http.Client{})
	symbols, err := sut.GetAllSymbols()
	assert.Nil(t, err)

	assert.NotEmpty(t, symbols)
}

func TestNasdaqProvider_GetPriceHistory_ShouldNotBeEmpty(t *testing.T) {
	sut := New(http.Client{})
	history, err := sut.GetPriceHistory("AAPL")
	assert.Nil(t, err)

	assert.NotEmpty(t, history)
	assert.NotEmpty(t, history[0].Open)
	assert.NotEmpty(t, history[0].High)
	assert.NotEmpty(t, history[0].Low)
	assert.NotEmpty(t, history[0].Close)
	assert.NotEmpty(t, history[0].Volume)
}

func TestNasdaqProvider_GetPriceHistory_GivenBadSymbol_ShouldError(t *testing.T) {
	sut := New(http.Client{})
	_, err := sut.GetPriceHistory("asdfasdfvxcv")
	assert.NotNil(t, err)
}

func TestNasdaqProvider_GetCompanyInfo_ShouldNotBeEmpty(t *testing.T) {
	sut := New(http.Client{})
	company, err := sut.GetCompanyInfo("AAPL")
	assert.Nil(t, err)

	assert.NotEmpty(t, company.Symbol)
	assert.NotEmpty(t, company.Sector)
	assert.NotEmpty(t, company.Industry)
	assert.NotEmpty(t, company.CompanyName)
	assert.NotEmpty(t, company.Exchange)
}