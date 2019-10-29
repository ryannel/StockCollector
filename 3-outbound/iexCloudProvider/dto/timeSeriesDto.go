package dto

type StockPriceSnapshotDto struct {
	Date           string  `json:"date"`           // Formatted as yyyy-MM-dd
	Open           float64 `json:"open"`           // Adjusted data for historical dates. Split adjusted only.
	Close          float64 `json:"close"`          // Adjusted data for historical dates. Split adjusted only.
	High           float64 `json:"high"`           // Adjusted data for historical dates. Split adjusted only.
	Low            float64 `json:"low"`            // Adjusted data for historical dates. Split adjusted only.
	Volume         int     `json:"volume"`         // Adjusted data for historical dates. Split adjusted only.
	UOpen          int     `json:"uOpen"`          // Adjusted data for historical dates.
	UClose         float64 `json:"uClose"`         // Adjusted data for historical dates.
	UHigh          float64 `json:"uHigh"`          // Adjusted data for historical dates.
	ULow           float64 `json:"uLow"`           // Adjusted data for historical dates.
	UVolume        int     `json:"uVolume"`        // Adjusted data for historical dates.
	Change         int     `json:"change"`         // Adjusted data for historical dates.
	ChangePercent  int     `json:"changePercent"`  // Change percent from previous trading day.
	Label          string  `json:"label"`          // A human readable format of the date depending on the range.
	ChangeOverTime int     `json:"changeOverTime"` // Change from previous trading day.
}
