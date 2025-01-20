package cmd

import (
	"cli-database/database"
	"fmt"
)

func begin(args []string, db *database.Database) (string, error) {
	if len(args) != 0 {
		err := fmt.Errorf("ERR BEGIN command do not receive arguments")
		return "", err
	}

	msg := db.BeginTransaction()
	return msg, nil
}

func rollback(args []string, db *database.Database) (string, error) {
	if len(args) != 0 {
		err := fmt.Errorf("ERR ROLLBACK command do not receive arguments")
		return "", err
	}

	msg := db.Rollback()
	return msg, nil
}

func commit(args []string, db *database.Database) (string, error) {
	if len(args) != 0 {
		err := fmt.Errorf("ERR COMMIT command do not receive arguments")
		return "", err
	}

	msg := db.Commit()
	return msg, nil
}
