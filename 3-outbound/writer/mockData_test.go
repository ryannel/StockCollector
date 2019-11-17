package writer

import (
	"stockCollector/2-core/outboundProviders"
	"time"
)

func getCompanyList(fixedTime time.Time) []outboundProviders.CompanyInfo {
	return []outboundProviders.CompanyInfo{
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
