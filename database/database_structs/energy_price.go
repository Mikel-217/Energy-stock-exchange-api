package databasestructs

import (
	"time"

	"mikel-kunze.com/energy-stock-exchange-api/database"
	"mikel-kunze.com/energy-stock-exchange-api/logging"
)

type EnergyPriceStruct struct {
	EnergyPriceId  uint // Primary key
	CurrentDate    time.Time
	BestTimeToBuy  uint // foreign key
	BestTimeToSell uint // foreign key
}

// Inserts the given EnergyPriceStruct into the database
func (e *EnergyPriceStruct) InsertIntoDatabase() bool {

	query := "INSERT INTO EnergyPrice() VALUES(DEFAULT, ?, ?, ?);"
	queryArgs := []any{e.CurrentDate, e.BestTimeToBuy, e.BestTimeToSell}

	if result := database.ExecuteSQL(query, queryArgs); result.ErrorMsg != nil {
		logging.Log(logging.Error, result.ErrorMsg.Error())
		return false
	}
	return true
}

func (e *EnergyPriceStruct) UpdateBestTimes(bestTimeToById, bestTimeToSellId uint) bool {

	query := "UPDATE EnergyPrice SET BestTimeToBuy = ?, BestTimeToSell = ? WHERE EnergyPriceId = ?;"
	queryArgs := []any{bestTimeToById, bestTimeToSellId, e.EnergyPriceId}

	if result := database.ExecuteSQL(query, queryArgs); result.ErrorMsg != nil {
		logging.Log(logging.Error, result.ErrorMsg.Error())
		return false
	}

	return true
}

type DateAndPriceStruct struct {
	DatePriceId   uint
	Date          time.Time
	Price         float32
	EnergyPriceId uint
}

// Inserts the given DataAndPriceStruct
func (d *DateAndPriceStruct) InsertIntoDatabase() *database.Result {

	query := "INSERT INTO DateAndPrice() VALUES(DEFAULT, ?, ?, ?);"
	queryArgs := []any{d.Date, d.Price, d.EnergyPriceId}

	result := database.ExecuteSQL(query, queryArgs)

	if result.ErrorMsg != nil {
		logging.Log(logging.Error, result.ErrorMsg.Error())
	}

	return result
}
