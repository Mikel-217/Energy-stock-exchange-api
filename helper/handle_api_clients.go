package helper

import (
	"fmt"

	apiclient "mikel-kunze.com/energy-stock-exchange-api/api_client"
	"mikel-kunze.com/energy-stock-exchange-api/structs"
)

// TODO:
// - Add a func which handels the responses from all clients
// - Add a func which stores everything in the db

func HandleAllClients(clienst []*apiclient.ApiClientStruct, responseChan chan structs.EnergyPriceStruct) {

	// we start the fetching from all clients
	for i := range clienst {

		// every client gets there own goroutine
		go clienst[i].StartFetchingData()
	}

	for {
		// for the future use select -> if more then one chan
		data := <-responseChan
		go processResponse(data)
	}
}

func processResponse(e structs.EnergyPriceStruct) {

	firstInsert := true // to indicate if its the first time trying or not

retry:
	// first we insert the given energy struct empty
	givenEnergyStruct := e.ConvertToDatabaseStruct()

	ok, resultParent := givenEnergyStruct.InsertIntoDatabase()

	if !ok && firstInsert {
		fmt.Println("Failed to insert. Please check logs. Retrying... ")
		firstInsert = false
		goto retry
	}

	firstInsert = true

retryBuy:

	// then we insert the best buy time
	givenBestTimeToBuy := e.BestTimeToBuy.ConvertToDatabaseStruct()
	givenBestTimeToBuy.EnergyPriceId = resultParent.LastId

	ok, resultBuy := givenBestTimeToBuy.InsertIntoDatabase()

	if !ok && firstInsert {
		fmt.Println("Failed to insert. Please check logs. Retrying... ")
		firstInsert = false
		goto retryBuy
	}

	// set the id to the parent struct
	givenEnergyStruct.BestTimeToBuy = resultBuy.LastId

	firstInsert = true
retrySell:

	// then we insert the best time to sell
	givenBestTimeToSell := e.BestTimeToSell.ConvertToDatabaseStruct()
	givenBestTimeToSell.EnergyPriceId = resultParent.LastId

	ok, resultSell := givenBestTimeToSell.InsertIntoDatabase()

	if !ok && firstInsert {
		fmt.Println("Failed to insert. Please check logs. Retrying... ")
		firstInsert = false
		goto retrySell
	}

	// set the id to the parent struct
	givenEnergyStruct.BestTimeToSell = resultSell.LastId

	// Updates the given struct
	if !givenEnergyStruct.UpdateBestTimes(resultBuy.LastId, resultSell.LastId) {
		fmt.Println("Failed to update. Please check database!")
	}
}
