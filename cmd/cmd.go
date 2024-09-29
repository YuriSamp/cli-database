package cmd

import (
	"cli-database/database"
	"fmt"
	"strings"
)

func Execute(input []string, db *database.Database) error {
	command := strings.ToUpper(input[0])
	args := input[1:]

	switch command {
	case "SET":
	case "GET":
	case "BEGIN":
		return begin(args, db)
	case "ROLLBACK":
		return rollback(args, db)
	case "COMMIT":
	default:
		return fmt.Errorf("ERR unknown command: %s", command)
	}

	fmt.Print(db)
	return nil
}

func begin(args []string, db *database.Database) error {
	if len(args) != 0 {
		return fmt.Errorf("ERR This command do not receive arguments")
	}

	db.BeginTransaction()
	return nil
}

func rollback(args []string, db *database.Database) error {
	if len(args) != 0 {
		return fmt.Errorf("ERR This command do not receive arguments")
	}

	return db.Rollback()
}
