package main

import (
	"fmt"
	"net/http"
	"os"

	"mikel-kunze.com/energy-stock-exchange-api/handlers"
	"mikel-kunze.com/energy-stock-exchange-api/helper"
	"mikel-kunze.com/energy-stock-exchange-api/logging"
	"mikel-kunze.com/energy-stock-exchange-api/startup"
)

// TODO:
// - implement server console to manage stuff easyer

func main() {

	// Create the tables -> if it fails we cannot continue
	if !startup.CreateDatabaseTables() {
		// so whoever reads this, can finally feel like a Ferrari F1 driver
		fmt.Println("Check the logs and your database connection! \n Program cannot start, but we are checking...")
		os.Exit(1)
	}

	// builds all api clients and then handels them in another goroutine
	allClients, responseChan := helper.BuildAllClients()
	go helper.HandleAllClients(allClients, responseChan)

	fmt.Println("Setup successful, now starting web-server...")

	mux := http.NewServeMux()

	// our api interfaces
	mux.HandleFunc("GET /price", handlers.HandlePriceRequests)
	mux.HandleFunc("GET /recommendation", handlers.HandleRecommendationRequests)

	// starts the server and checks for a error
	if err := http.ListenAndServe(":8080", mux); err != nil {
		logging.Log(logging.Error, err.Error())
		fmt.Println("Check the logs. Cannot start the server...")
		os.Exit(1)
	}
}
