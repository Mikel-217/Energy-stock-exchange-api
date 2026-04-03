package databasestructs

import (
	"time"

	"mikel-kunze.com/energy-stock-exchange-api/database"
	"mikel-kunze.com/energy-stock-exchange-api/logging"
)

// ============ EnergyPriceStruct ============

type EnergyPriceStruct struct {
	EnergyPriceId  uint      `json:"energy-price-id"` // Primary key
	CurrentDate    time.Time `json:"date"`
	BestTimeToBuy  uint      `json:"best-time-to-buy"`  // foreign key
	BestTimeToSell uint      `json:"best-time-to-sell"` // foreign key
}

// Inserts the given EnergyPriceStruct into the database
func (e *EnergyPriceStruct) InsertIntoDatabase() (bool, *database.Result) {

	query := "INSERT INTO EnergyPrice() VALUES(DEFAULT, ?, DEFAULT, DEFAULT);"
	queryArgs := []any{e.CurrentDate}

	result := database.ExecuteSQL(query, queryArgs)

	if result.ErrorMsg != nil {

		logging.Log(logging.Error, result.ErrorMsg.Error())
		return false, result
	}

	return true, result
}

// Updates the given EnergyPricestruct with new ids
func (e *EnergyPriceStruct) UpdateBestTimes(bestTimeToById, bestTimeToSellId uint) bool {

	query := "UPDATE EnergyPrice SET BestTimeToBuy = ?, BestTimeToSell = ? WHERE EnergyPriceId = ?;"
	queryArgs := []any{bestTimeToById, bestTimeToSellId, e.EnergyPriceId}

	if result := database.ExecuteSQL(query, queryArgs); result.ErrorMsg != nil {
		logging.Log(logging.Error, result.ErrorMsg.Error())
		return false
	}

	return true
}

// ============ DateAndPriceStruct ============

type DateAndPriceStruct struct {
	DatePriceId   uint      `json:"date-price-id"`
	Date          time.Time `json:"time"`
	Price         float32   `json:"price"`
	EnergyPriceId uint      `json:"parent-id"`
}

// Inserts the given DataAndPriceStruct
func (d *DateAndPriceStruct) InsertIntoDatabase() (bool, *database.Result) {

	query := "INSERT INTO DateAndPrice() VALUES(DEFAULT, ?, ?, ?);"
	queryArgs := []any{d.Date, d.Price, d.EnergyPriceId}

	result := database.ExecuteSQL(query, queryArgs)

	if result.ErrorMsg != nil {
		logging.Log(logging.Error, result.ErrorMsg.Error())
		return false, result
	}

	return true, result
}
