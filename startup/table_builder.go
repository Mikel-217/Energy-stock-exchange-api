package startup

import (
	"encoding/json"
	"fmt"
	"os"

	"mikel-kunze.com/energy-stock-exchange-api/database"
	"mikel-kunze.com/energy-stock-exchange-api/logging"
)

type jsonData struct {
	TableName    string `json:"Table-Name"`
	TableCommand string `json:"Table-Command"`
}

// Creates the database tables
// Returns if successfull or not
func CreateDatabaseTables() bool {

	// we check if the app was started before -> else we setup the tables
	if firstStartup := os.Getenv("Started"); firstStartup != "" {
		return true
	}

	currAppPath, _ := os.Getwd()
	path := currAppPath + "/startup/json/tables.json"

	// get all file content
	fileContent, err := os.ReadFile(path)

	if err != nil {
		logging.Log(logging.Error, err.Error())
		return false
	}

	jsonData := make([]jsonData, 10)

	if err := json.Unmarshal(fileContent, &jsonData); err != nil {
		logging.Log(logging.Error, err.Error())
		return false
	}

	// loop over all entrys
	for i := range jsonData {

		// excecutes the query and checks for errors
		if result := database.ExecuteSQL(jsonData[i].TableCommand, nil); result.ErrorMsg != nil {
			logging.Log(logging.Error, result.ErrorMsg.Error())
			return false
		}
	}

	// FIXME
	os.Setenv("Started", "true")

	fmt.Println("Created tables successful")
	return true
}
