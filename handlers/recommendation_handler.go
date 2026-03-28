package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"mikel-kunze.com/energy-stock-exchange-api/database"
	databasestructs "mikel-kunze.com/energy-stock-exchange-api/database/database_structs"
	"mikel-kunze.com/energy-stock-exchange-api/logging"
)

// TODO:
// implement a handler that handels:
// - the current best time to buy and sell
// - the best time to buy and sell for a date  ?date=yyyy-mm-dd
// a date range ?start=yyyy-mm-dd&end=yyyy-mm-dd
func HandleRecommendationRequests(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	switch {
	case q.Has("date"):
	case q.Has("start"):
	default:
		result, ok := handleGivenDate(time.Now().String())

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(result)
		w.WriteHeader(http.StatusOK)
	}
}

// handels the request for a given date
func handleGivenDate(date string) ([]byte, bool) {

	givenDate, err := time.Parse("2006-01-02 15:04:05", date)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return []byte{}, false
	}

	// creates our db read builder
	dbReadBuilder := database.CreateNewBuilder[databasestructs.EnergyPriceStruct]()
	// sets everything up
	db := dbReadBuilder.AddQuery("SELECT * FROM EnergyPrice WHERE Date = ?;").
		AddQueryParams([]any{givenDate}).
		Build()

	result := db.GetData()

	if len(result) == 0 {
		return []byte{}, false
	}

	jsonData, err := json.Marshal(result)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return []byte{}, false
	}

	return jsonData, true
}
