package database

import (
	"reflect"

	"mikel-kunze.com/energy-stock-exchange-api/logging"
)

type ReadDataClient[T any] struct {
	Query       string
	QueryParams []any
}

// To read data from the database
func (r *ReadDataClient[T]) GetData() []T {

	var results []T

	// Creates our database connection
	db := CreateDBConn()

	if db == nil {
		return results
	}

	defer db.Close()

	rows, err := db.Query(r.Query, r.QueryParams...)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return results
	}

	columns, _ := rows.Columns()

	for rows.Next() {

		var data T

		v := reflect.ValueOf(&data).Elem()

		fieldPointers := make([]any, len(columns))

		for i := range columns {
			if i < v.NumField() {
				fieldPointers[i] = v.Field(i).Addr().Interface()
			}
		}

		if err := rows.Scan(fieldPointers...); err != nil {
			logging.Log(logging.Error, err.Error())
			break
		}

		results = append(results, data)
	}

	return results
}
