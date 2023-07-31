package main

import (
	"fmt"
	"log"
)

func dsn(db string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, db)
}

func checkError(errMsg string, err error) {
	if err != nil {
		fmt.Println(errMsg)
		log.Fatal(err)
	}
}

/*
func createContext() context.Context {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	return ctx
}*/
