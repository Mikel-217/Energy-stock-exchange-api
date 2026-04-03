package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"mikel-kunze.com/energy-stock-exchange-api/database"
	databasestructs "mikel-kunze.com/energy-stock-exchange-api/database/database_structs"
	"mikel-kunze.com/energy-stock-exchange-api/logging"
)

// serves all prices for a date ?all=year-month-day
// serve prices for a date range ?start=year-month-day1&end=year-month-day2
func HandlePriceRequests(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	switch {
	case q.Has("all"):
		date := q.Get("all")

		if date == "" {
			date = time.Now().Format("2006-01-02 15:04:05")
		}

		response, ok := handleAllDataPrice(date)

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There was an error getting your data!"))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return

	case q.Has("start"):
		date1 := q.Get("start")
		date2 := q.Get("end")

		if date1 == "" || date2 == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No dates given"))
			return
		}

		response, ok := handleTimeStampPrice(date1, date2)

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
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
			w.Write([]byte("No id found"))
			return
		}

		response, ok := handlePriceId(id)

		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There was an error getting your data"))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// handels the current date
// returns the json data and an indication of success
func handleAllDataPrice(date string) ([]byte, bool) {

	splitedDate := strings.Split(date, "-")

	if len(splitedDate) == 0 {
		return []byte{}, false
	}

	// converts our date to midnight
	dateMidNight := time.Date(convertToInt(splitedDate[0]), time.Month(convertToInt(splitedDate[1])), convertToInt(splitedDate[2]), 0, 0, 0, 0, time.UTC)

	// creates our db read builder
	dbReadBuilder := database.CreateNewBuilder[databasestructs.DateAndPriceStruct]()
	// sets everything up
	db := dbReadBuilder.AddQuery("SELECT * FROM DateAndPrice WHERE Date >= ? AND Date < ? + INTERVAL 1 DAY;").
		AddQueryParams([]any{dateMidNight, dateMidNight}).
		Build()

	// gets the data
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

// handels the date range data
// returns the json data and an indication of success
func handleTimeStampPrice(start, end string) ([]byte, bool) {

	startDate := strings.Split(start, "-")

	// checks if the lengh is null
	if len(startDate) == 0 {
		return []byte{}, false
	}

	dateStart := time.Date(convertToInt(startDate[0]), time.Month(convertToInt(startDate[1])), convertToInt(startDate[2]), 0, 0, 0, 0, time.UTC)

	endDate := strings.Split(end, "-")

	if len(endDate) == 0 {
		return []byte{}, false
	}

	dateEnd := time.Date(convertToInt(endDate[0]), time.Month(convertToInt(endDate[1])), convertToInt(endDate[2]), 0, 0, 0, 0, time.UTC)

	// creates our sql read builder
	dbReadBuilder := database.CreateNewBuilder[databasestructs.DateAndPriceStruct]()

	// builds our client
	db := dbReadBuilder.AddQuery("SELECT * FROM DateAndPrice WHERE Date >= ? AND Date < ?;").
		AddQueryParams([]any{dateStart, dateEnd}).
		Build()

	// excecutes the query
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

func convertToInt(d string) int {

	i, _ := strconv.Atoi(d)
	return i
}

// handels the price by id
func handlePriceId(id int) ([]byte, bool) {

	dbBuilder := database.CreateNewBuilder[databasestructs.DateAndPriceStruct]()

	db := dbBuilder.AddQuery("SELECT * FROM DateAndPrice WHERE DatePriceId = ?;").AddQueryParams([]any{id}).Build()

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
