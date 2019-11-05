package nasdaqProvider

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"io/ioutil"
	"net/http"
	"stockCollector/2-core/outboundProviders"
	"stockCollector/infrastructure/helpers"
	"strings"
	"time"
)



func New(httpClient http.Client) NasdaqProvider {
	return NasdaqProvider{
		baseUrl: "https://api.nasdaq.com/api/",
		httpClient:     httpClient,
	}
}

type NasdaqProvider struct {
	httpClient     http.Client
	baseUrl string
}

func (provider *NasdaqProvider) GetCompanyInfo(symbol string) (outboundProviders.Company, error) {
	urlBuilder, err := helpers.NewUrlBuilder(provider.baseUrl)
	if err != nil {
		return outboundProviders.Company{}, fmt.Errorf("error creating URL builder for nasdaq provider - %s : %w", provider.baseUrl, err)
	}
	urlBuilder.AppendPath(fmt.Sprintf("company/%s/company-profile", symbol))

	requestUrl := urlBuilder.GetUrl()
	responseBody, err := provider.httpGet(requestUrl)
	if err != nil {
		return outboundProviders.Company{}, err
	}

	var responseData CompanyInfoDto
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return outboundProviders.Company{}, fmt.Errorf("unable to parse json httpResponse from Nasdaq - Error: %w", err)
	}

	return mapCompanyInfo(responseData), nil
}

func (provider *NasdaqProvider) GetPriceHistory(symbol string) ([]outboundProviders.StockPriceSnapshot, error) {
	urlBuilder, err := helpers.NewUrlBuilder(provider.baseUrl)
	if err != nil {
		return nil, fmt.Errorf("error creating URL builder for nasdaq provider - %s : %w", provider.baseUrl, err)
	}
	urlBuilder.AppendPath(fmt.Sprintf("quote/%s/historical", symbol))
	urlBuilder.AddQueryParameter("assetclass", "stocks")
	urlBuilder.AddQueryParameter("fromdate", "1800-01-01")
	urlBuilder.AddQueryParameter("limit", "10000000")
	urlBuilder.AddQueryParameter("todate", time.Now().Format("2006-01-02"))

	requestUrl := urlBuilder.GetUrl()
	responseBody, err := provider.httpGet(requestUrl)
	if err != nil {
		return nil, err
	}

	var responseData HistoricalPriceDto
	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, fmt.Errorf("unable to parse json httpResponse from Nasdaq - Error: %w", err)
	}

	return mapStockPriceSnapshots(responseData)
}

func (provider *NasdaqProvider) httpGet(requestUrl string) ([]byte, error) {
	httpResponse, err := provider.httpClient.Get(requestUrl)
	if err != nil {
		return nil, fmt.Errorf("nasdaq error sending request - %s: %w", requestUrl, err)
	}
	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nasdaq service error %s: %s", requestUrl, httpResponse.Status)
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read nasdaq httpResponse body - Error : %w", err)
	}

	return responseBody, nil
}

func (provider *NasdaqProvider) GetAllSymbols() ([]string, error) {
	listedSymbols, err := getListedSymbols()
	if err != nil {
		return nil, err
	}

	unlistedSymbols, err := getUnlistedListedSymbols()
	if err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(listedSymbols) + len(unlistedSymbols))
	for key := range listedSymbols {
		keys = append(keys, key)
	}
	for key := range unlistedSymbols {
		keys = append(keys, key)
	}

	return keys, nil
}

func getListedSymbols() (map[string]nasdaqListedSymbolsFtpDto, error) {
	listedResponse, err := getFtpFile("symboldirectory/nasdaqlisted.txt")
	if err != nil {
		return nil, err
	}

	rows, err := readCsv(listedResponse)
	if err != nil {
		return nil, err
	}

	symbols := make(map[string]nasdaqListedSymbolsFtpDto)
	for _, row := range rows {
		symbols[row[0]] = nasdaqListedSymbolsFtpDto{
			Symbol:          row[0],
			SecurityName:    row[1],
			MarketCategory:  row[2],
			TestIssue:       row[3],
			FinancialStatus: row[4],
			RoundLotSize:    row[5],
			Etf:             row[6],
			NextShares:      row[7],
		}
	}

	return symbols, nil
}

func getUnlistedListedSymbols() (map[string]nasdaqUnlistedSymbolsFtpDto, error) {
	listedResponse, err := getFtpFile("symboldirectory/otherlisted.txt")
	if err != nil {
		return nil, err
	}

	rows, err := readCsvRowWise(listedResponse)
	if err != nil {
		return nil, err
	}

	symbols := make(map[string]nasdaqUnlistedSymbolsFtpDto)
	for _, row := range rows {
		symbols[row[0]] = nasdaqUnlistedSymbolsFtpDto{
			ActSymbol:    row[0],
			SecurityName: row[1],
			Exchange:     row[2],
			CqsSymbol:    row[3],
			Etf:          row[4],
			RoundLotSize: row[5],
			TestIssue:    row[6],
			NasdaqSymbol: row[7],
		}
	}

	return symbols, nil
}

func getFtpFile(path string) (*ftp.Response, error) {
	client, err := ftp.Dial("ftp.nasdaqtrader.com:21")
	if err != nil {
		return nil, err
	}

	err = client.Login("anonymous", "anonymous@domain.com")
	if err != nil {
		return nil, err
	}

	return client.Retr(path)
}

func readCsv(fileReader io.Reader) ([][]string, error) {
	csvr := csv.NewReader(fileReader)
	csvr.Comma = '|'
	result, err := csvr.ReadAll()
	return result, err
}

func readCsvRowWise(fileReader io.Reader) ([][]string, error) {
	csvr := csv.NewReader(fileReader)
	csvr.Comma = '|'

	// Skip header row
	_, err := csvr.Read()
	if err != nil {
		return nil, err
	}

	var csvRows [][]string
	for {
		row, err := csvr.Read()
		// Last line of file has wrong structure
		if err != nil && strings.HasSuffix(err.Error(), "wrong number of fields") {
			err = nil
			break
		}
		if err != nil {
			return nil, err
		}

		// End of file
		if strings.HasPrefix(row[0], "File Creation Time") {
			break
		}

		csvRows = append(csvRows, row)
	}

	return csvRows, nil
}
