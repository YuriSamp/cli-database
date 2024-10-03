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
	case "DEL":
		return delete(args, db)
	case "HELP":
		help()
		return nil
	default:
		return fmt.Errorf("%s is not a command. See help to list commands", command)
	}
}

func get(args []string, db *database.Database) error {
	if len(args) != 1 {
		return fmt.Errorf("ERR GET <key> - Syntax error")
	}

	key := args[0]
	msg := db.Get(key)

	fmt.Println(msg)

	return nil
}

func set(args []string, db *database.Database) error {
	if len(args) != 2 {
		return fmt.Errorf("ERR SET <key> - <value> - Syntax error")
	}

	key := args[0]
	value := args[1]

	msg := db.Set(key, value)

	fmt.Println(msg)

	return nil
}

func delete(args []string, db *database.Database) error {
	if len(args) != 1 {
		return fmt.Errorf("ERR DEL <key> - Syntax error")
	}

	key := args[0]

	msg := db.Delete(key)

	fmt.Println(msg)

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

func help() {
	fmt.Println("usage: command <arg1> <arg2>")
	fmt.Println("   commands:")
	fmt.Println("     GET Receive one argument and retrieve a value")
	fmt.Println("     SET Receive two arguments and set a value")
	fmt.Println("     BEGIN don't receive arguments, init a transaction")
	fmt.Println("     ROLLBACK don't receive arguments, rollback a transaction")
	fmt.Println("     COMMIT don't receive arguments, commit a transaction")
}
