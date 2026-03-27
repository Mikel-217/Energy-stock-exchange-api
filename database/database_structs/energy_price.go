package databasestructs

import "time"

type EnergyPriceStruct struct {
	EnergyPriceId  uint // Primary key
	CurrentDate    time.Time
	BestTimeToBuy  uint // foreign key
	BestTimeToSell uint // foreign key
}

func (e *EnergyPriceStruct) InsertIntoDatabase() {
	// TODO
}

type DateAndPriceStruct struct {
	DatePriceId   uint
	Date          time.Time
	Price         float32
	EnergyPriceId uint
}

func (d *DateAndPriceStruct) InsertIntoDatabase() {
	// TODO
}
