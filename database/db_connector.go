package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"mikel-kunze.com/energy-stock-exchange-api/logging"
)

// Creates a new database connection
func CreateDBConn() *sql.DB {

	db, err := sql.Open("mysql", os.Getenv("DB-Conn"))

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return nil
	}

	return db
}
