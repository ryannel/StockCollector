package nasdaqProvider

import (
	"fmt"
	"math"
	"stockCollector/2-core/outboundProviders"
	"strconv"
	"strings"
	"time"
)

func mapStockPriceSnapshots(src HistoricalPriceDto) ([]outboundProviders.StockPriceSnapshot, error) {
	 results := make([]outboundProviders.StockPriceSnapshot, src.Data.TotalRecords)

	for i, price := range src.Data.TradesTable.Rows {
		dateTime, err := time.Parse("01/02/2006", price.Date)
		if err != nil {
			return nil, fmt.Errorf("unable to parse Nasdaq price histroy - datetime: %w", err)
		}

		open, err := cleanFloatStrings(price.Open)
		if err != nil {
			return nil, fmt.Errorf("unable to parse Nasdaq price - open: %w", err)
		}

		high, err := cleanFloatStrings(price.High)
		if err != nil {
			return nil, fmt.Errorf("unable to parse Nasdaq price - high: %w", err)
		}

		low, err := cleanFloatStrings(price.Low,)
		if err != nil {
			return nil, fmt.Errorf("unable to parse Nasdaq price - low: %w", err)
		}

		closePrice, err := cleanFloatStrings(price.Close)
		if err != nil {
			return nil, fmt.Errorf("unable to parse Nasdaq price - close: %w", err)
		}

		volumeFloat, err := cleanFloatStrings(price.Volume)
		if err != nil {
			return nil, fmt.Errorf("unable to parse Nasdaq price - volume: %w", err)
		}
		volume := int(math.Round(volumeFloat))

		results[i] = outboundProviders.StockPriceSnapshot{
			DateTime: dateTime,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    closePrice,
			Volume:   volume,
		}
	}

	return results, nil
}

func cleanFloatStrings(value string) (float64, error) {
	cleanString := strings.ReplaceAll(value, "$", "")
	cleanString = strings.ReplaceAll(cleanString, ",", "")
	cleanString = strings.ReplaceAll(cleanString, " ", "")
	cleanString = strings.ReplaceAll(cleanString, "N/A", "0")
	result, err := strconv.ParseFloat(cleanString, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse Nasdaq float while cleaning - %s : %w", value, err)
	}
	return result, err
}

func mapCompanyInfo(src CompanyInfoDto) outboundProviders.CompanyInfo {
	data := src.Data

	return outboundProviders.CompanyInfo{
		CompanyName: data.CompanyName.Value,
		Industry:     data.Industry.Value,
		Sector:       data.Sector.Value,
		Symbol:       data.Symbol.Value,
		Exchange:     "NASDAQ",
	}
}

func mapInsiderActivity(src InsiderActivityDto) ([]outboundProviders.InsiderActivity, error) {
	var result []outboundProviders.InsiderActivity

	for _, item := range src.Data.TransactionTable.Rows {
		lastDate, err := time.Parse("01/02/2006", item.LastDate)
		if err != nil {
			return nil, fmt.Errorf("unable to parse insider activity last date: %w", err)
		}

		result = append(result, outboundProviders.InsiderActivity{
			Insider:         item.Insider,
			Relation:        item.Relation,
			LastDate:        lastDate,
			TransactionType: item.TransactionType,
			OwnType:         item.OwnType,
			SharesTraded:    item.SharesTraded,
			LastPrice:       item.LastPrice,
			SharesHeld:      item.SharesHeld,
		})
	}

	return result, nil
}