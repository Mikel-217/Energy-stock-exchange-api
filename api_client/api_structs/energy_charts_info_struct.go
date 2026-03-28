package apistructs

import (
	"slices"
	"time"

	"mikel-kunze.com/energy-stock-exchange-api/structs"
)

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
func (e *EnergyChartsApiStruct) ConvertToEnergyStruct() structs.EnergyPriceStruct {

	convertedStructs := make([]structs.DateAndPriceStruct, 100)

	for i := range e.UnixSeconds {

		var newConvert structs.DateAndPriceStruct

		newConvert.Time = time.Unix(int64(e.UnixSeconds[i]), 0)
		newConvert.Price = e.Price[i]

		convertedStructs = append(convertedStructs, newConvert)
	}

	return structs.EnergyPriceStruct{
		Date:                   time.Now(),
		AllPricesAndTheyreTime: convertedStructs,
		BestTimeToBuy: slices.MinFunc(convertedStructs, func(a, b structs.DateAndPriceStruct) int {
			if a.Price < b.Price {
				return -1
			} else if a.Price > b.Price {
				return 1
			}
			return 0
		}),
		BestTimeToSell: slices.MaxFunc(convertedStructs, func(a, b structs.DateAndPriceStruct) int {
			if a.Price < b.Price {
				return -1
			} else if a.Price > b.Price {
				return 1
			}
			return 0
		}),
	}
}
