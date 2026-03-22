package structs

import (
	"time"
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

// The time and the price to that given time
type DateAndPriceStruct struct {
	Time  time.Time
	Price float32
}
