package models

import "time"

type StockPriceSnapshot struct {
	Id        uint `gorm:"primary_key"`
	CompanyId uint
	DateTime  time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    int
}

type Company struct {
	Id           uint `gorm:"primary_key"`
	CompanyName  string
	Industry     string
	Sector       string
	Symbol       string
	Exchange     string
	Cusip        string
	PriceHistory []StockPriceSnapshot
}
