package iexCloudProvider

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

// Use sandbox env for all integration tests
func TestMain(m *testing.M) {
	baseUrl = "https://sandbox.iexapis.com/stable"
	os.Exit(m.Run())
}

func Test_GetDailyStockPriceHistory1y_shouldNotError(t *testing.T) {
	sut, err := New("Tpk_3bd8245273f744bda84217fa121e5a0f", http.Client{})
	assert.Nil(t, err)

	result, err := sut.GetDailyStockPriceHistory1y("AAPL")
	assert.Nil(t, err)

	assert.NotEmpty(t, result)
}

func Test_GetCompanyInfo_shouldNotError(t *testing.T) {
	sut, err := New("Tpk_3bd8245273f744bda84217fa121e5a0f", http.Client{})
	assert.Nil(t, err)

	result, err := sut.GetCompanyInfo("AAPL")
	assert.Nil(t, err)

	assert.NotEqual(t, "", result.Exchange)
	assert.NotEqual(t, "", result.Symbol)
	assert.NotEqual(t, "", result.Sector)
	assert.NotEqual(t, "", result.Industry)
	assert.NotEqual(t, "", result.CompanyName)
}
