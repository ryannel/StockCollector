package outboundProviders

import (
	"time"
)

type StockProviderInterface interface {
	GetDailyStockPriceHistory20y(symbol string) ([]StockPriceSnapshot, error)
	GetDailyStockPriceHistory20d(symbol string) ([]StockPriceSnapshot, error)
}

type StockPriceSnapshot struct {
	DateTime time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int
}

type Company struct {
	CompanyName  string
	Industry     string
	Sector       string
	Symbol       string
	Exchange     string
	Cusip        string
	PriceHistory []StockPriceSnapshot
}
