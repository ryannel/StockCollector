package writer

import (
	"bufio"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"stockCollector/2-core/outboundProviders"
	"testing"
	"time"
)

func TestLinesWriter_ShouldWriteCompany(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err)

	outputFilePath := filepath.Join(wd, "testCompany.jsonl")

	sut, err := NewJsonLinesWriter(outputFilePath)
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

	err = sut.AppendLine(insertCmp)
	assert.Nil(t, err)

	jsonByte, err := ioutil.ReadFile(outputFilePath)
	assert.Nil(t, err)

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

	err = os.Remove(outputFilePath)
	assert.Nil(t, err)
}

func TestLinesWriter_ShouldWriteCompanies(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err)

	outputFilePath := filepath.Join(wd, "testCompanies.jsonl")
	_ = os.Remove(outputFilePath)

	sut, err := NewJsonLinesWriter(outputFilePath)
	assert.Nil(t, err)

	fixedTime, err := time.Parse("2006-01-02", "2019-01-02")
	assert.Nil(t, err)

	insertCompanies := getCompanyList(fixedTime)

	for _, company := range insertCompanies {
		err = sut.AppendLine(company)
		assert.Nil(t, err)
	}

	file, err := os.Open(outputFilePath)
	assert.Nil(t, err)

	reader := bufio.NewReader(file)
	var resultCompanies []outboundProviders.Company
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			err = nil // End of file
			break
		}
		if err != nil {
			break
		}

		var resultCompany outboundProviders.Company
		err = json.Unmarshal(line, &resultCompany)
		assert.Nil(t, err)

		resultCompanies = append(resultCompanies, resultCompany)
	}
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

	err = file.Close()
	assert.Nil(t, err)
	err = os.Remove(outputFilePath)
	assert.Nil(t, err)
}

func TestLinesWriter_ShouldReadCompanies(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err)

	outputFilePath := filepath.Join(wd, "testReadCompanies.jsonl")
	_ = os.Remove(outputFilePath)

	sut, err := NewJsonLinesWriter(outputFilePath)
	assert.Nil(t, err)

	fixedTime, err := time.Parse("2006-01-02", "2019-01-02")
	assert.Nil(t, err)

	insertCompanies := getCompanyList(fixedTime)

	for _, company := range insertCompanies {
		err = sut.AppendLine(company)
		assert.Nil(t, err)
	}

	resultCompanies, err := sut.ReadCompanies()
	assert.Nil(t, err)

	assert.NotEmpty(t, resultCompanies)
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

	err = os.Remove(outputFilePath)
	assert.Nil(t, err)
}
