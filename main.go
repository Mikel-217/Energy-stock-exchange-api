package main

import (
	"fmt"
	"net/http"
	"os"

	apiclient "mikel-kunze.com/energy-stock-exchange-api/api_client"
	"mikel-kunze.com/energy-stock-exchange-api/logging"
	"mikel-kunze.com/energy-stock-exchange-api/startup"
)

// save all api clients here
var AllApiClients []apiclient.ApiClientStruct

func main() {

	// Create the tables -> if it fails we cannot continue
	if !startup.CreateDatabaseTables() {
		// so whoever reads this, can finally feel like a Ferrari F1 driver
		fmt.Println("Check the logs and your database connection! \n Program cannot start, but we are checking...")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	// starts the server and checks for a error
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logging.Log(logging.Error, err.Error())
		os.Exit(1)
	}
}
