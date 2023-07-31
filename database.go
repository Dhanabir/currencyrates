package main

import (
	"database/sql"
	"time"
)

func dbConnection() (*sql.DB, error) {
	db, err := sql.Open(sqlDriver, dsn(""))
	checkError("1", err)
	//defer db.Close()
	_, err = db.Exec("create database if not exists " + dbname)
	checkError("2", err)
	db.Close()

	//create new connection with dbName
	db, err = sql.Open(sqlDriver, dsn(dbname))
	checkError("3", err)
	//defer db.Close()
	err = db.Ping()
	checkError("4", err)

	return db, err
}

func createTable(db *sql.DB) error {
	query := `Create table if not exists rates(id int not null primary key auto_increment, Dates DATE NOT NULL UNIQUE, EUR FLOAT NOT NULL default 1.00, USD FLOAT NOT NULL, GBP FLOAT NOT NULL)`
	_, err := db.Exec(query)
	checkError("5", err)
	return err
}

/*
func insertValues(db *sql.DB, rsp responseBody) error {
	query := "insert into rates(Dates, USD, GBP) values(?, ?, ?)"
	//insert, err := db.Exec("insert into rates(EUR, GBP) values(1.8, 0.5)")
	stmnt, err := db.Prepare(query)
	checkError("6", err)
	defer stmnt.Close()
	res, err := stmnt.Exec(rsp.Date, rsp.Rates.USD, rsp.Rates.GBP)
	checkError("7", err)
	_ = res
	return err
}
*/

func todaysDataExists(db *sql.DB) (bool, error) {
	var n int
	query := "select 1 from rates where Dates='" + time.Now().Format("2006-01-02") + "'"
	err := db.QueryRow(query).Scan(&n)
	if err == nil {
		return true, nil
	} else if err == sql.ErrNoRows {
		return false, nil
	}
	return false, err
}

func insertValue(db *sql.DB, rsp latestRatesBody) error {
	query := "insert into rates(Dates, USD, GBP) values(?, ?, ?)"
	//insert, err := db.Exec("insert into rates(EUR, GBP) values(1.8, 0.5)")
	stmnt, err := db.Prepare(query)
	checkError("6", err)
	defer stmnt.Close()
	res, err := stmnt.Exec(rsp.Date, rsp.Rates.USD, rsp.Rates.GBP)
	checkError("7", err)
	_ = res
	return err
}

func updateValue(db *sql.DB, rsp latestRatesBody) error {
	query := "update rates set USD=?, GBP=? where Dates =?"
	//insert, err := db.Exec("insert into rates(EUR, GBP) values(1.8, 0.5)")
	stmnt, err := db.Prepare(query)
	checkError("11", err)
	defer stmnt.Close()
	res, err := stmnt.Exec(rsp.Rates.USD, rsp.Rates.GBP, rsp.Date)
	checkError("12", err)
	_ = res
	return err
}

/*
stmt, err := db.Prepare("select name from users where id = ?")
if err != nil {
	log.Fatal(err)
}
defer stmt.Close()
var name string
err = stmt.QueryRow(1).Scan(&name)
if err != nil {
	log.Fatal(err)
}
fmt.Println(name)



var (
	id int
	name string
)
rows, err := db.Query("select id, name from users where id = ?", 1)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(&id, &name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id, name)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}

*/
