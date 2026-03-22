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

// This func needs the complete SQL-Statement as a string and it arguments
// It returns an struct with the last Id and an error to indicate success
func ExecuteSQL(sqlQuery string, args []any) *Result {

	db := CreateDBConn()

	if db == nil {
		return &Result{
			LastId:   0,
			ErrorMsg: errors.New("DB error"),
		}
	}

	defer db.Close()

	// Excecute the given command with all arguments
	queryResult, err := db.Exec(sqlQuery, args...)

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
