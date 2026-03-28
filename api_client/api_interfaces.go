package apiclient

import "mikel-kunze.com/energy-stock-exchange-api/structs"

// To convert any struct to our custom struct
type EnergyConverter interface {
	ConvertToEnergyStruct() structs.EnergyPriceStruct
}
