package iexCloudProvider

type StockPriceSnapshotDto struct {
	Date           string  `json:"date"`           // Formatted as yyyy-MM-dd
	Open           float64 `json:"open"`           // Adjusted data for historical dates. Split adjusted only.
	Close          float64 `json:"close"`          // Adjusted data for historical dates. Split adjusted only.
	High           float64 `json:"high"`           // Adjusted data for historical dates. Split adjusted only.
	Low            float64 `json:"low"`            // Adjusted data for historical dates. Split adjusted only.
	Volume         int     `json:"volume"`         // Adjusted data for historical dates. Split adjusted only.
	UOpen          float64 `json:"uOpen"`          // Adjusted data for historical dates.
	UClose         float64 `json:"uClose"`         // Adjusted data for historical dates.
	UHigh          float64 `json:"uHigh"`          // Adjusted data for historical dates.
	ULow           float64 `json:"uLow"`           // Adjusted data for historical dates.
	UVolume        int     `json:"uVolume"`        // Adjusted data for historical dates.
	Change         float64 `json:"change"`         // Adjusted data for historical dates.
	ChangePercent  float64 `json:"changePercent"`  // Change percent from previous trading day.
	Label          string  `json:"label"`          // A human readable format of the date depending on the range.
	ChangeOverTime float64 `json:"changeOverTime"` // Change from previous trading day.
}

type CompanyDto struct {
	Symbol         string   `json:"symbol"`
	CompanyName    string   `json:"companyName"`
	Exchange       string   `json:"exchange"`
	Industry       string   `json:"industry"`
	Website        string   `json:"website"`
	Description    string   `json:"description"`
	CEO            string   `json:"CEO"`
	SecurityName   string   `json:"securityName"`
	IssueType      string   `json:"issueType"`
	Sector         string   `json:"sector"`
	PrimarySicCode int      `json:"primarySicCode"`
	Employees      int      `json:"employees"`
	Tags           []string `json:"tags"`
	Address        string   `json:"address"`
	Address2       string   `json:"address2"`
	State          string   `json:"state"`
	City           string   `json:"city"`
	Zip            string   `json:"zip"`
	Country        string   `json:"country"`
	Phone          string   `json:"phone"`
}
