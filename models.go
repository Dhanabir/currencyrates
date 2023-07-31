package main

const (
	sqlDriver    = "mysql"
	username     = "root"
	password     = "root"
	hostname     = "127.0.0.1:3306"
	dbname       = "currencyrates"
	url          = "http://api.exchangeratesapi.io/v1/"
	access_key   = "3d25927a1e78dcd2d9b620531cee52f7"
	baseCurrency = "EUR"
	currencies   = "USD, GBP"
)

type latestRatesBody struct {
	Success      bool       `json:"success"`
	Time         int        `json:"timestamp"`
	BaseCurrency string     `json:"base"`
	Date         string     `json:"date"`
	Rates        latesRates `json:"rates"`
}

/*
type HistoricRatesBody struct {
	Success      bool   `json:"success"`
	Timeseries   bool   `json:"timestamp"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	BaseCurrency string `json:"base"`
	Date         string `json:"date"`
	Rates        rates  `json:"rates"`
}
*/

type latesRates struct {
	USD float64 `json:"USD"`
	GBP float64 `json:"GBP"`
}

/*
type historicRates struct {
}*/
