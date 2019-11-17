package outboundProviders

import (
	"time"
)

type StockProviderInterface interface {
	GetCompanyInfo(symbol string) (CompanyInfo, error)
	GetStockPriceHistory(symbol string) ([]StockPriceSnapshot, error)
	GetInsiderActivity(symbol string) ([]InsiderActivity, error)
}

type CompanyInfo struct {
	CompanyName  string
	Industry     string
	Sector       string
	Symbol       string
	Exchange     string
}

type StockPriceSnapshot struct {
	DateTime time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int
}

type InsiderActivity struct {
	Insider         string
	Relation        string
	LastDate        time.Time
	TransactionType string
	OwnType         string
	SharesTraded    string
	LastPrice       string
	SharesHeld      string
}