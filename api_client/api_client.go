package apiclient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"time"

	"mikel-kunze.com/energy-stock-exchange-api/logging"
	"mikel-kunze.com/energy-stock-exchange-api/structs"
)

type ApiClientStruct struct {
	Name    string // the name of the site
	BaseUrl string // the basic url
	FullUrl string // url with query

	RequiresKey bool   // if the API needs a key
	ApiKey      string // if it needs a key -> use this field

	GetInterval time.Duration // the interval we want to fetch the data
	StructType  any           // the given struct for the api

	SendBackChan chan<- structs.EnergyPriceStruct // to send data back
	ctx          context.Context                  // so we can cancel it
}

// Starts fetching data in the given interval
// Sends it back into a chan
func (a *ApiClientStruct) StartFetchingData() {

	ticker := time.NewTicker(a.GetInterval)

	defer ticker.Stop()

	if a.RequiresKey {
		// TODO: build the header
	}

	for {
		select {
		case <-a.ctx.Done():
			return
		case <-ticker.C:

			// sends a get request with the full url
			response, err := http.Get(a.FullUrl)

			if err != nil {
				logging.Log(logging.Error, err.Error())
				ticker.Reset(a.GetInterval)
				break
			}

			body, err := io.ReadAll(response.Body)

			// Gets the given struct type
			t := reflect.TypeOf(a.StructType)
			// Creates a new Struct from the given struct type
			givenStruct := reflect.New(t)

			// unmarshals it into our given struct
			if err := json.Unmarshal(body, givenStruct.Interface()); err != nil {
				logging.Log(logging.Error, err.Error())
				ticker.Reset(a.GetInterval)
				break
			}

			// We then check if it has the interface implemented and then transform our data
			if converter, ok := givenStruct.Elem().Interface().(EnergyConverter); ok {
				a.SendBackChan <- converter.ConvertToEnergyStruct()
			}
			ticker.Reset(a.GetInterval)
		}
	}
}
