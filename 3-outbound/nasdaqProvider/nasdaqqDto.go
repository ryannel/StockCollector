package nasdaqProvider

type HistoricalPriceDto struct {
	Data    SymbolData `json:"data"`
	Message string     `json:"message"`
	Status  Status     `json:"status"`
}

type Headers struct {
	Date   string `json:"date"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
}

type Rows struct {
	Date   string `json:"date"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
}

type TradesTable struct {
	Headers Headers `json:"headers"`
	Rows    []Rows  `json:"rows"`
}

type SymbolData struct {
	Symbol       string      `json:"symbol"`
	TotalRecords int         `json:"totalRecords"`
	TradesTable  TradesTable `json:"tradesTable"`
}

type Status struct {
	RCode            int         `json:"rCode"`
	BCodeMessage     interface{} `json:"bCodeMessage"`
	DeveloperMessage interface{} `json:"developerMessage"`
}

type CompanyInfoDto struct {
	Data    CompanyData `json:"data"`
	Message string      `json:"message"`
	Status  Status      `json:"status"`
}

type KeyValue struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type CompanyData struct {
	ModuleTitle        KeyValue `json:"ModuleTitle"`
	CompanyName        KeyValue `json:"CompanyName"`
	Symbol             KeyValue `json:"Symbol"`
	Address            KeyValue `json:"Address"`
	Phone              KeyValue `json:"Phone"`
	Industry           KeyValue `json:"Industry"`
	Sector             KeyValue `json:"Sector"`
	Region             KeyValue `json:"Region"`
	CompanyDescription KeyValue `json:"CompanyDescription"`
}

// ===========================
// FTP Data Transfer Objects
// ===========================

type nasdaqListedSymbolsFtpDto struct {
	Symbol          string
	SecurityName    string
	MarketCategory  string
	TestIssue       string
	FinancialStatus string
	RoundLotSize    string
	Etf             string
	NextShares      string
}

type nasdaqUnlistedSymbolsFtpDto struct {
	ActSymbol    string
	SecurityName string
	Exchange     string
	CqsSymbol    string
	Etf          string
	RoundLotSize string
	TestIssue    string
	NasdaqSymbol string
}
