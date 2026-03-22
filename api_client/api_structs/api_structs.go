package apistructs

import "mikel-kunze.com/energy-stock-exchange-api/structs"

// ========== NOTE ==========
// Set ur own structs here

// struct for the "api.energy-charts.info" api
type EnergyChartsApiStruct struct {
	LicenseInfo string    `json:"license_info"`
	UnixSeconds []int     `json:"unix_seconds"`
	Price       []float32 `json:"price"`
	Unit        string    `json:"unit"`
	Deprecated  bool      `json:"deprecated"`
}

// Converts the EnergychartsApiStruct to the custom EnergyPriceStruct
func (e *EnergyChartsApiStruct) ConvertToEnergyPriceStruct() structs.EnergyPriceStruct {

	// TODO: Convert the unix seconds to time.time And also set the given prices to the right time
	return structs.EnergyPriceStruct{}
}
