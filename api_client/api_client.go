package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
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

	if a.GetInterval == 0 {
		a.GetInterval = time.Duration(24 * time.Hour)
	}

	ticker := time.NewTicker(a.GetInterval)

	defer ticker.Stop()

	httpRequest, err := http.NewRequest("GET", a.FullUrl, nil)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		fmt.Println("Error starting request in: ", a.Name, ". \n Check logs for detailed information")
		return
	}

	if a.RequiresKey {
		httpRequest.Header.Add("Authorization", a.ApiKey)
	}

	for {
		select {
		case <-a.ctx.Done():
			return
		case <-ticker.C:

			// sends a get request with the full url
			response, err := http.DefaultClient.Do(httpRequest)

			// if there is an response err, we just reset and continue with our day
			if err != nil {
				logging.Log(logging.Error, err.Error())
				ticker.Reset(a.GetInterval)
				break
			}

			body, err := io.ReadAll(response.Body)

			// checks for errors from the body
			if err != nil {
				logging.Log(logging.Error, err.Error())
				break
			}

			// gets the selected struct type
			t := reflect.TypeOf(a.StructType)

			// if the type is nil -> just return
			if t == nil {
				logging.Log(logging.Error, fmt.Sprintf("No struct type defined for client: %s", a.Name))
				fmt.Println("No struct type defined for client: ", a.Name)
				break
			}

			// Creates a new Struct from the given struct type
			givenStruct := reflect.New(t).Interface()

			// unmarshals it into our given struct
			if err := json.Unmarshal(body, &givenStruct); err != nil {
				logging.Log(logging.Error, err.Error())
				ticker.Reset(a.GetInterval)
				break
			}

			// We then check if it has the interface implemented and then transform our data
			if converter, ok := givenStruct.(EnergyConverter); ok {
				// sends the data back to our chan
				a.SendBackChan <- converter.ConvertToEnergyStruct()
			} else {
				// logs the error
				logging.Log(logging.Error, fmt.Sprintf("No interface for: %s", a.Name))
				fmt.Println("Interface isnt there for: ", a.Name)
			}

			ticker.Reset(a.GetInterval)
		}
	}
}
