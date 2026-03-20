package main

import (
	"fmt"
	"net/http"
	"os"

	"mikel-kunze.com/energy-stock-exchange-api/logging"
	"mikel-kunze.com/energy-stock-exchange-api/startup"
)

func main() {

	// Create the tables -> if it fails we cannot continue
	if !startup.CreateDatabaseTables() {
		fmt.Println("Check the logs and your database connection! \n Program cannot start :(")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	// starts the server and checks for a error
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logging.Log(logging.Error, err.Error())
		os.Exit(1)
	}
}
