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
		return set(args, db)
	case "GET":
		return get(args, db)
	case "BEGIN":
		return begin(args, db)
	case "ROLLBACK":
		return rollback(args, db)
	case "COMMIT":
		return commit(args, db)
	default:
		return fmt.Errorf("ERR unknown command: %s", command)
	}
}

func get(args []string, db *database.Database) error {
	if len(args) != 1 {
		return fmt.Errorf("ERR GET <key> - Syntax error")
	}

	key := args[0]
	db.Get(key)

	return nil
}

func set(args []string, db *database.Database) error {
	if len(args) != 2 {
		return fmt.Errorf("ERR SET <key> - <value> - Syntax error")
	}

	key := args[0]
	value := args[1]

	db.Set(key, value)

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

func commit(args []string, db *database.Database) error {
	if len(args) != 0 {
		return fmt.Errorf("ERR This command do not receive arguments")
	}

	return db.Commit()
}
