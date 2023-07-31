package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func getCurrencyRates(requestValue string) latestRatesBody {
	res, err := http.Get(url + "/" + requestValue + "?access_key=" + access_key + "&base=" + baseCurrency + "&symbols=" + currencies)
	checkError("100", err)

	responseData, err := io.ReadAll(res.Body)
	checkError("200", err)

	//var response map[string]interface{}
	//fmt.Println(string(responseData))
	var response latestRatesBody
	err = json.Unmarshal(responseData, &response)
	checkError("300", err)

	return response
}

/*
func getHistoricRates() {
	res, err := http.Get(url + "/timeseries?access_key=" + access_key + "&base=" + baseCurrency + "&symbols=" + currencies + "&start_date=" + time.Now().Format("2006-01-02") + "&end_date=" + time.Now().AddDate(0, 0, -10).Format("2006-01-02"))
	checkerr("101", err)

	responseData, err := io.ReadAll(res.Body)
	checkError("201", err)


}*/
