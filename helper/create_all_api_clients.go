package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	apiclient "mikel-kunze.com/energy-stock-exchange-api/api_client"
	apistructs "mikel-kunze.com/energy-stock-exchange-api/api_client/api_structs"
	"mikel-kunze.com/energy-stock-exchange-api/logging"
	"mikel-kunze.com/energy-stock-exchange-api/structs"
)

type ClientJson struct {
	Name        string `json:"api-name"`
	BaseUrl     string `json:"api-base-url"`
	FullUrl     string `json:"api-full-url"`
	RequiresKey bool   `json:"api-requires-key"`
	ApiKey      string `json:"api-key-site"`
	GetInterval string `json:"api-get-interval"`
	StructName  string `json:"api-struct-type"`
}

// ========== !NOTE! ==========
// Add new structs here
var structRegistry = map[string]any{
	"EnergyCharts": apistructs.EnergyChartsApiStruct{},
	// Add new structs here as you create them
}

// Builds all our api clients and the response chan
func BuildAllClients() ([]*apiclient.ApiClientStruct, chan structs.EnergyPriceStruct) {

	allClientData := getJsonData() // gets all the required data from the json config

	if len(allClientData) == 0 {
		fmt.Println("Please check the logs, there was an error building the clients. \n Program cannot start, but we are checking...")
		os.Exit(1)
	}

	ctx := context.Background()
	givenChan := make(chan structs.EnergyPriceStruct)

	allBuildClients := make([]*apiclient.ApiClientStruct, 0, len(allClientData))

	for i := range allClientData {
		timeDuration, _ := time.ParseDuration(allClientData[i].GetInterval)

		// creates our client builder
		clientBuilder := apiclient.NewApiClientBuilder()

		// builds our client
		client := clientBuilder.
			WithName(allClientData[i].Name).SetBaseUrl(allClientData[i].BaseUrl).SetFullUrl(allClientData[i].FullUrl).
			WithApiKey(allClientData[i].RequiresKey, allClientData[i].ApiKey).
			WithSructTyp(getStructType(allClientData[i].StructName)).
			SetInterval(timeDuration).SetCtx(ctx).SetOutputChan(givenChan).Build()

		// adds the client to the slice
		allBuildClients = append(allBuildClients, client)
	}

	return allBuildClients, givenChan
}

// gets all json data from the given file
func getJsonData() []ClientJson {
	currDir, _ := os.Getwd()

	filePath := path.Join(currDir, "helper", "json", "clientData.json")

	fileData, err := os.ReadFile(filePath)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return []ClientJson{}
	}

	allClients := make([]ClientJson, 0)

	if err := json.Unmarshal(fileData, &allClients); err != nil {
		logging.Log(logging.Error, err.Error())
		return []ClientJson{}
	}

	return allClients
}

// gets the given struct for the name
func getStructType(name string) any {
	val, ok := structRegistry[name]
	if !ok {
		logging.Log(logging.Error, "Unknown API name in config: "+name)
		return nil
	}
	return val
}
