package nasdaqProvider

import (
	"encoding/csv"
	"github.com/jlaffaye/ftp"
	"io"
	"strings"
)

type nasdaqListedSymbols struct {
	Symbol          string
	SecurityName    string
	MarketCategory  string
	TestIssue       string
	FinancialStatus string
	RoundLotSize    string
	Etf             string
	NextShares      string
}

type nasdaqUnlistedSymbols struct {
	ActSymbol    string
	SecurityName string
	Exchange     string
	CqsSymbol    string
	Etf          string
	RoundLotSize string
	TestIssue    string
	NasdaqSymbol string
}

func GetSymbolList() ([]string, error) {
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

func getListedSymbols() (map[string]nasdaqListedSymbols, error) {
	listedResponse, err := getFtpFile("symboldirectory/nasdaqlisted.txt")
	if err != nil {
		return nil, err
	}

	rows, err := readCsv(listedResponse)
	if err != nil {
		return nil, err
	}

	symbols := make(map[string]nasdaqListedSymbols)
	for _, row := range rows {
		symbols[row[0]] = nasdaqListedSymbols{
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

func getUnlistedListedSymbols() (map[string]nasdaqUnlistedSymbols, error) {
	listedResponse, err := getFtpFile("symboldirectory/otherlisted.txt")
	if err != nil {
		return nil, err
	}

	rows, err := readCsvRowWise(listedResponse)
	if err != nil {
		return nil, err
	}

	symbols := make(map[string]nasdaqUnlistedSymbols)
	for _, row := range rows {
		symbols[row[0]] = nasdaqUnlistedSymbols{
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
