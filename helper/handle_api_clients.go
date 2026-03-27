package helper

import (
	apiclient "mikel-kunze.com/energy-stock-exchange-api/api_client"
	"mikel-kunze.com/energy-stock-exchange-api/structs"
)

// TODO:
// - Add a func which handels the responses from all clients
// - Add a func which stores everything in the db

func HandleAllClients(clienst []apiclient.ApiClientStruct, responseChan chan<- structs.EnergyPriceStruct) {
	// TODO
}
