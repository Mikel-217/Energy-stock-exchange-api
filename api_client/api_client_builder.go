package apiclient

import (
	"context"
	"time"

	"mikel-kunze.com/energy-stock-exchange-api/structs"
)

type ApiClientBuilderStruct struct {
	client ApiClientStruct
}

// Creates a new API-Client builder
func NewApiClientBuilder() *ApiClientBuilderStruct {
	return &ApiClientBuilderStruct{
		client: ApiClientStruct{
			Name:        "Default",
			GetInterval: 30 * time.Second,
			RequiresKey: false,
		},
	}
}

// Sets the name
func (a *ApiClientBuilderStruct) WithName(name string) *ApiClientBuilderStruct {
	a.client.Name = name
	return a
}

// Sets the base url
func (a *ApiClientBuilderStruct) SetBaseUrl(url string) *ApiClientBuilderStruct {
	a.client.BaseUrl = url
	return a
}

// Sets the full url
func (a *ApiClientBuilderStruct) SetFullUrl(url string) *ApiClientBuilderStruct {
	a.client.FullUrl = url
	return a
}

// Sets the interval
func (a *ApiClientBuilderStruct) SetInterval(interval time.Duration) *ApiClientBuilderStruct {
	a.client.GetInterval = interval
	return a
}

// Sets the wanted struct
func (a *ApiClientBuilderStruct) WithSructTyp(structType any) *ApiClientBuilderStruct {
	a.client.StructType = structType
	return a
}

// Sets the wanted API key -> if there is one
func (a *ApiClientBuilderStruct) WithApiKey(needsKey bool, apiKey string) *ApiClientBuilderStruct {
	a.client.RequiresKey = true
	a.client.ApiKey = apiKey
	return a
}

// Sets the Output chan
func (a *ApiClientBuilderStruct) SetOutputChan(givenChan chan<- structs.EnergyPriceStruct) *ApiClientBuilderStruct {
	a.client.SendBackChan = givenChan
	return a
}

// Adds context
func (a *ApiClientBuilderStruct) SetCtx(ctx context.Context) *ApiClientBuilderStruct {
	a.client.ctx = ctx
	return a
}

// Builds the api client
func (a *ApiClientBuilderStruct) Build() *ApiClientStruct {
	return &a.client
}
