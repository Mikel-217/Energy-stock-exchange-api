package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"mikel-kunze.com/energy-stock-exchange-api/database"
	databasestructs "mikel-kunze.com/energy-stock-exchange-api/database/database_structs"
	"mikel-kunze.com/energy-stock-exchange-api/logging"
)

// Gets the data for recommendations
// - the current best time to buy and sell
// - the best time to buy and sell for a date  ?date=yyyy-mm-dd
// a date range ?start=yyyy-mm-dd&end=yyyy-mm-dd
func HandleRecommendationRequests(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	switch {
	case q.Has("date"):
		date := q.Get("date")

		if date == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Date required!"))
			return
		}

		givenDate, err := time.Parse("2006-01-02", date)

		if err != nil {
			logging.Log(logging.Error, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error parsing date"))
		}

		result, ok := handleGivenDate(givenDate.Format("2006-01-02"))

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("No data for: " + givenDate.String()))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	case q.Has("start"):

		start := q.Get("start")
		end := q.Get("end")

		if start == "" || end == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Dates are null!"))
			return
		}

		response, ok := handleTimeSpanRecommendation(start, end)

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There was an error getting data"))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return

	case q.Has("id"):

		id, err := strconv.Atoi(q.Get("id"))

		if err != nil {
			logging.Log(logging.Error, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No Id given"))
		}

		response, ok := handleGivenId(id)

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There was no data"))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	default:
		result, ok := handleGivenDate(time.Now().Format("2006-01-02"))

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

// handels the request for a given date
func handleGivenDate(date string) ([]byte, bool) {

	startTime := date + " 00:00:00"
	endTime := date + " 23:59:59"

	// creates our db read builder
	dbReadBuilder := database.CreateNewBuilder[databasestructs.EnergyPriceStruct]()
	// sets everything up
	db := dbReadBuilder.AddQuery("SELECT * FROM EnergyPrice WHERE CurrentDate >= ? AND CurrentDate <= ? ;").
		AddQueryParams([]any{startTime, endTime}).
		Build()

	result := db.GetData()

	if len(result) == 0 {
		logging.Log(logging.Information, "Nothing in the database for: recommendation_handler")
		return []byte{}, false
	}

	jsonData, err := json.Marshal(result)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return []byte{}, false
	}

	return jsonData, true
}

// handels the request for a given timespan
func handleTimeSpanRecommendation(start, end string) ([]byte, bool) {

	firstDate := start + " 00:00:00"
	secondDate := end + " 23:59:59"

	dbReader := database.CreateNewBuilder[databasestructs.EnergyPriceStruct]()

	db := dbReader.AddQuery("SELECT * FROM EnergyPrice WHERE CurrentDate >= ? AND CurrentDate <= ?;").
		AddQueryParams([]any{firstDate, secondDate}).
		Build()

	results := db.GetData()

	if len(results) == 0 {
		logging.Log(logging.Information, "No data")
		return []byte{}, false
	}

	jsonData, err := json.Marshal(results)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return []byte{}, false
	}

	return jsonData, true
}

// handels the request for ids
func handleGivenId(id int) ([]byte, bool) {

	dbBuilder := database.CreateNewBuilder[databasestructs.EnergyPriceStruct]()

	db := dbBuilder.AddQuery("SELECT * FROM EnergyPrice WHERE EnergyPriceId = ?").AddQueryParams([]any{id}).Build()

	data := db.GetData()

	if len(data) == 0 {
		return []byte{}, false
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return []byte{}, false
	}

	return jsonData, true
}
