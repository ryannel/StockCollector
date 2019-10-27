package dto

import "time"

type Iso8601 struct {
	time.Time
}

func (iso *Iso8601) UnmarshalJSON(data []byte) error {
	timeString := string(data)
	resultTime, err := time.Parse(`"`+time.RFC3339Nano+`"`, timeString)
	if err != nil {
		resultTime, err = time.Parse(`"`+"2006-01-02"+`"`, timeString)
		if err != nil {
			return err
		}
	}
	iso.Time = resultTime
	return nil
}

type TimeSeriesDailyDto struct {
	MetaData        MetaDataDto                       `json:"Meta Data"`
	TimeSeriesDaily map[Iso8601]StockPriceSnapshotDto `json:"Time Series (Daily)"`
	Error           string                            `json:"Error Message"`
	Information     string                            `json:"Information"`
}

type StockPriceSnapshotDto struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type MetaDataDto struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}
