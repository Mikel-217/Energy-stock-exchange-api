package database

import (
	"errors"

	"mikel-kunze.com/energy-stock-exchange-api/logging"
)

// NOTE:
// Only use this for executing DELETE, UPDATE or INSERT commads!!

type Result struct {
	LastId   uint
	ErrorMsg error
}

// This func needs the complete SQL-Statement as a string and it arguments as an string slice
// It returns an struct with the last Id and an error to indicate success
func ExecuteSQL(sqlQuery string, args []string) *Result {

	arguments := make([]interface{}, len(args))

	// Converts the given strings to an interface
	for i, arg := range args {
		arguments[i] = arg
	}

	db := CreateDBConn()

	if db == nil {
		return &Result{
			LastId:   0,
			ErrorMsg: errors.New("DB error"),
		}
	}

	defer db.Close()

	queryResult, err := db.Exec(sqlQuery, arguments...)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return &Result{
			LastId:   0,
			ErrorMsg: err,
		}
	}

	Id, _ := queryResult.LastInsertId()

	return &Result{
		LastId:   uint(Id),
		ErrorMsg: nil,
	}
}
