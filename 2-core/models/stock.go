package models

type Stock struct {
	Name     string
	Exchange string
	Symbol   string
	Sector   string
	Industry string
	Cusips   []Cusip
}

type Cusip struct {
	Cusip string
}
