package structs

import (
	"time"

	databasestructs "mikel-kunze.com/energy-stock-exchange-api/database/database_structs"
)

// ======== NOTE =========
// These structs are just used for the API or our internal stuff

// To display all data and the best times to buy / sell
type EnergyPriceStruct struct {
	Date                   time.Time // for which date is this data
	AllPricesAndTheyreTime []DateAndPriceStruct
	BestTimeToBuy          DateAndPriceStruct
	BestTimeToSell         DateAndPriceStruct
}

// Transforms the given struct to our database struct
func (e *EnergyPriceStruct) ConvertToDatabaseStruct() *databasestructs.EnergyPriceStruct {

	databaseST := new(databasestructs.EnergyPriceStruct)
	databaseST.CurrentDate = e.Date

	return databaseST
}

// The time and the price to that given time
type DateAndPriceStruct struct {
	Time  time.Time
	Price float32
}

// Transforms the given struct to our database struct
func (d *DateAndPriceStruct) ConvertToDatabaseStruct() *databasestructs.DateAndPriceStruct {

	databaseST := new(databasestructs.DateAndPriceStruct)
	databaseST.Date = d.Time
	databaseST.Price = d.Price

	return databaseST
}
