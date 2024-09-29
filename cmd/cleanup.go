package cmd

import (
	"cli-database/database"
	"encoding/json"
	"fmt"
	"os"
)

func CleanUp(db *database.Database) {
	firstLayer := db.GetFirstLayer()

	body, err := json.Marshal(firstLayer)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	os.WriteFile("./database.json", body, 0666)
	os.Exit(0)
}
