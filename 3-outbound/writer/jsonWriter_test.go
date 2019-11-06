package writer

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"stockCollector/2-core/outboundProviders"
	"testing"
	"time"
)

func TestWriter_ShouldWriteCompany(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err)

	sut, err := New(wd)
	assert.Nil(t, err)

	fixedTime, err := time.Parse("2006-01-02", "2019-01-02")
	assert.Nil(t, err)

	insertCmp := outboundProviders.Company{
		CompanyName:  "companyName",
		Industry:     "industry",
		Sector:       "sector",
		Symbol:       "symbol",
		Exchange:     "exchange",
		Cusip:        "",
		PriceHistory: []outboundProviders.StockPriceSnapshot{
			{
				DateTime: fixedTime,
				Open:     200,
				High:     500,
				Low:      300,
				Close:    250,
				Volume:   5000,
			},
			{
				DateTime: fixedTime.AddDate(0, 1, 12),
				Open:     500.025,
				High:     700.50,
				Low:      300.6,
				Close:    200.2,
				Volume:   3500,
			},
		},
	}

	outputFile := "testCompany.json"
	err = sut.WriteJson(insertCmp, outputFile)
	assert.Nil(t, err)

	jsonByte, _ := ioutil.ReadFile(outputFile)
	var resultCompany outboundProviders.Company
	err = json.Unmarshal(jsonByte, &resultCompany)
	assert.Nil(t, err)

	assert.Equal(t, "companyName", resultCompany.CompanyName)
	assert.Equal(t, "sector", resultCompany.Sector)
	assert.Equal(t, "symbol", resultCompany.Symbol)
	assert.Equal(t, "exchange", resultCompany.Exchange)

	assert.Equal(t, 2, len(resultCompany.PriceHistory))
	for i, snapshot := range resultCompany.PriceHistory {
		if snapshot.DateTime == fixedTime {
			assert.Equal(t, fixedTime, resultCompany.PriceHistory[i].DateTime)
			assert.Equal(t, float64(200), resultCompany.PriceHistory[i].Open)
			assert.Equal(t, float64(500), resultCompany.PriceHistory[i].High)
			assert.Equal(t, float64(300), resultCompany.PriceHistory[i].Low)
			assert.Equal(t, float64(250), resultCompany.PriceHistory[i].Close)
			assert.Equal(t, 5000, resultCompany.PriceHistory[i].Volume)
		} else {
			assert.Equal(t, fixedTime.AddDate(0, 1, 12), resultCompany.PriceHistory[i].DateTime)
			assert.Equal(t, float64(500.025), resultCompany.PriceHistory[i].Open)
			assert.Equal(t, float64(700.50), resultCompany.PriceHistory[i].High)
			assert.Equal(t, float64(300.6), resultCompany.PriceHistory[i].Low)
			assert.Equal(t, float64(200.2), resultCompany.PriceHistory[i].Close)
			assert.Equal(t, 3500, resultCompany.PriceHistory[i].Volume)
		}
	}

	err = os.Remove(outputFile)
	assert.Nil(t, err)
}

func TestWriter_ShouldWriteCompanies(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err)

	sut, err := New(wd)
	assert.Nil(t, err)

	fixedTime, err := time.Parse("2006-01-02", "2019-01-02")
	assert.Nil(t, err)

	insertCompanies := getCompanyList(fixedTime)

	outputFile := "testCompanies.json"
	err = sut.WriteJson(insertCompanies, outputFile)
	assert.Nil(t, err)

	jsonByte, _ := ioutil.ReadFile(outputFile)
	var resultCompanies []outboundProviders.Company
	err = json.Unmarshal(jsonByte, &resultCompanies)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(resultCompanies))

	for i, company := range resultCompanies {
		if company.CompanyName == "companyName" {
			assert.Equal(t, "companyName", resultCompanies[i].CompanyName)
			assert.Equal(t, "sector", resultCompanies[i].Sector)
			assert.Equal(t, "symbol", resultCompanies[i].Symbol)
			assert.Equal(t, "exchange", resultCompanies[i].Exchange)
			assert.Equal(t, 2, len(resultCompanies[i].PriceHistory))

		} else {
			assert.Equal(t, "companyName2", resultCompanies[i].CompanyName)
			assert.Equal(t, "sector2", resultCompanies[i].Sector)
			assert.Equal(t, "symbol2", resultCompanies[i].Symbol)
			assert.Equal(t, "exchange2", resultCompanies[i].Exchange)
			assert.Equal(t, 2, len(resultCompanies[i].PriceHistory))
		}
	}

	err = os.Remove(outputFile)
	assert.Nil(t, err)
}

func TestJsonWriter_ReadJson(t *testing.T) {
	fixedTime, err := time.Parse("2006-01-02", "2019-01-02")
	assert.Nil(t, err)

	wd, err := os.Getwd()
	assert.Nil(t, err)

	sut, err := New(wd)
	assert.Nil(t, err)

	insertCompanies := getCompanyList(fixedTime)


	outputFile := "readerTest.json"
	err = sut.WriteJson(insertCompanies, outputFile)
	assert.Nil(t, err)

	var resultCompanies []outboundProviders.Company
	err = sut.ReadJson("readerTest.json", &resultCompanies)
	assert.Nil(t, err)

	for i, company := range resultCompanies {
		if company.CompanyName == "companyName" {
			assert.Equal(t, "companyName", resultCompanies[i].CompanyName)
			assert.Equal(t, "sector", resultCompanies[i].Sector)
			assert.Equal(t, "symbol", resultCompanies[i].Symbol)
			assert.Equal(t, "exchange", resultCompanies[i].Exchange)
			assert.Equal(t, 2, len(resultCompanies[i].PriceHistory))

		} else {
			assert.Equal(t, "companyName2", resultCompanies[i].CompanyName)
			assert.Equal(t, "sector2", resultCompanies[i].Sector)
			assert.Equal(t, "symbol2", resultCompanies[i].Symbol)
			assert.Equal(t, "exchange2", resultCompanies[i].Exchange)
			assert.Equal(t, 2, len(resultCompanies[i].PriceHistory))
		}
	}

	err = os.Remove(outputFile)
	assert.Nil(t, err)
}

func getCompanyList(fixedTime time.Time) []outboundProviders.Company {
	return []outboundProviders.Company{
		{
			CompanyName:  "companyName",
			Industry:     "industry",
			Sector:       "sector",
			Symbol:       "symbol",
			Exchange:     "exchange",
			Cusip:        "",
			PriceHistory: []outboundProviders.StockPriceSnapshot{
				{
					DateTime: fixedTime,
					Open:     200,
					High:     500,
					Low:      300,
					Close:    250,
					Volume:   5000,
				},
				{
					DateTime: fixedTime.AddDate(0, 1, 12),
					Open:     500.025,
					High:     700.50,
					Low:      300.6,
					Close:    200.2,
					Volume:   3500,
				},
			},
		}, {
			CompanyName: "companyName2",
			Industry:    "industry2",
			Sector:      "sector2",
			Symbol:      "symbol2",
			Exchange:    "exchange2",
			Cusip:       "",
			PriceHistory: []outboundProviders.StockPriceSnapshot{
				{
					DateTime: fixedTime,
					Open:     2002,
					High:     5002,
					Low:      3002,
					Close:    2502,
					Volume:   50002,
				},
				{
					DateTime: fixedTime.AddDate(0, 1, 12),
					Open:     500.0252,
					High:     700.502,
					Low:      300.62,
					Close:    200.22,
					Volume:   35002,
				},
			},
		},
	}
}