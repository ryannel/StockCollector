package nasdaqProvider

import (
	"encoding/csv"
	"github.com/jlaffaye/ftp"
	"io"
	"strings"
)

func New() error {
	return nil
}

type Nasdaq struct {}

type SymbolDto struct {
	Symbol string
	SecurityName string
	MarketCategory string
	TestIssue string
	FinancialStatus string
	RoundLotSize string
	Etf string
	NextShares string
}

func (nasdaq *Nasdaq) GetSymbolList() ([]string, error) {
	listedResponse, err := nasdaq.getCsv("symboldirectory/nasdaqlisted.txt")
	if err != nil {
		return nil, err
	}

	listedSymbols, err := nasdaq.readCsv(listedResponse)

	otherListedResponse, err := nasdaq.getCsv("symboldirectory/otherlisted.txt")
	if err != nil {
		return nil, err
	}

	otherSymbols, err := nasdaq.readCsv(otherListedResponse)

	for key, value := range otherSymbols {
		listedSymbols[key] = value
	}

	keys := make([]string, 0, len(listedSymbols))
	for key := range listedSymbols {
		keys = append(keys, key)
	}

	return keys, nil
}

func (nasdaq *Nasdaq) getCsv(path string) (*ftp.Response, error) {
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

func (nasdaq *Nasdaq) readCsv(fileReader io.Reader) (map[string]SymbolDto, error)  {
	csvr := csv.NewReader(fileReader)
	csvr.Comma = '|'
	symbols := make(map[string]SymbolDto)

	// Skip header row
	_, err := csvr.Read()
	if err != nil {
		return nil, err
	}

	for {
		row, err := csvr.Read()
		if err != nil {
			return nil, err
		}

		// End of file
		if strings.HasPrefix(row[0], "File Creation Time:") {
			return symbols, nil
		}

		symbols[row[0]] = SymbolDto{
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
}