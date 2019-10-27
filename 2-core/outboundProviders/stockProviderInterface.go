package outboundProviders

import (
    "time"
)

type StockProviderInterface interface {
    GetDaily(symbol string) (StockPriceSeries, error)
    GetDailyRecent(symbol string) (StockPriceSeries, error)
}

type StockPriceSeries map[time.Time]StockPriceSnapshot

type StockPriceSnapshot struct {
    Open   float64
    High   float64
    Low    float64
    Close  float64
    Volume float64
}