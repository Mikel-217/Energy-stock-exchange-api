package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	apiclient "mikel-kunze.com/energy-stock-exchange-api/api_client"
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

func BuildAllClients() ([]apiclient.ApiClientStruct, chan<- structs.EnergyPriceStruct) {

	allClientData := getJsonData()

	if len(allClientData) == 0 {
		fmt.Println("Please check the loggs, there was an error building the clients. \n Program exists...")
		os.Exit(1)
	}

	ctx := context.Background()
	givenChan := make(chan<- structs.EnergyPriceStruct)

	allBuildClients := make([]apiclient.ApiClientStruct, 10)

	for i := range allClientData {
		timeDuration, _ := time.ParseDuration(allClientData[i].GetInterval)

		clientBuilder := apiclient.NewApiClientBuilder()

		client := clientBuilder.
			WithName(allClientData[i].Name).SetBaseUrl(allClientData[i].BaseUrl).SetFullUrl(allClientData[i].FullUrl).
			WithApiKey(allClientData[i].RequiresKey, allBuildClients[i].ApiKey).
			WithSructTyp(getStructType(allBuildClients[i].Name)).
			SetInterval(timeDuration).SetCtx(ctx).SetOutputChan(givenChan).Build()

		allBuildClients = append(allBuildClients, *client)
	}

	return allBuildClients, givenChan
}

// gets all json data from the given file
func getJsonData() []ClientJson {
	currDir, _ := os.Getwd()

	filePath := path.Join(currDir, "json", "clientData.json")

	fileData, err := os.ReadFile(filePath)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return []ClientJson{}
	}

	allClients := make([]ClientJson, 10)

	if err := json.Unmarshal(fileData, &allClients); err != nil {
		logging.Log(logging.Error, err.Error())
		return []ClientJson{}
	}

	return allClients
}

// gets the given struct for the name
func getStructType(name string) any {
	// TODO
	return nil
}
