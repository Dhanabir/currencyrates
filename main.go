package main

import (
	"database/sql"
	"time"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := dbConnection()
	checkError("0", err)
	defer db.Close()

	err = createTable(db)
	checkError("0.1", err)

	//fmt.Println("test ok")

	//check historic data first: top 10 then add
	//setHistoricRates(db)
	//add time constraint later on
	//setLatestRates(db)

}

func setHistoricRates(db *sql.DB) {
	for i := 1; i <= 10; i++ {
		response := getCurrencyRates(time.Now().AddDate(0, 0, -i).Format("2006-01-02"))
		err := insertValue(db, response)
		checkError("0.10.0", err)
	}
}

func setLatestRates(db *sql.DB) {
	exists, err := todaysDataExists(db)
	checkError("0.2", err)

	response := getCurrencyRates("latest")

	/*response := latestRatesBody{
		Date: "2023-07-28",
		Rates: rates{
			USD: 1.2,
			GBP: 0.8,
		},
	}*/

	if exists {
		err = updateValue(db, response)
		checkError("0.9", err)
	} else {
		err = insertValue(db, response)
		checkError("0.10", err)
	}
}
