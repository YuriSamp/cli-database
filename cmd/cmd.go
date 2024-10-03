package cmd

import (
	"cli-database/database"
	"fmt"
	"strings"
)

func Execute(input []string, db *database.Database) {
	command := strings.ToUpper(input[0])
	args := input[1:]

	switch command {
	case "SET":
		set(args, db)
	case "GET":
		get(args, db)
	case "BEGIN":
		begin(args, db)
	case "ROLLBACK":
		rollback(args, db)
	case "COMMIT":
		commit(args, db)
	case "DEL":
		delete(args, db)
	case "HELP":
		help()
	default:
		fmt.Printf("%s is not a command. See help to list commands \n", command)
	}
}

func get(args []string, db *database.Database) {
	if len(args) != 1 {
		fmt.Println("ERR GET <key> - Syntax error")
		return
	}

	key := args[0]
	msg := db.Get(key)

	fmt.Println(msg)
}

func set(args []string, db *database.Database) {
	if len(args) != 2 {
		fmt.Println("ERR SET <key> - <value> - Syntax error")
		return
	}

	key := args[0]
	value := args[1]

	msg := db.Set(key, value)
	fmt.Println(msg)
}

func delete(args []string, db *database.Database) {
	if len(args) != 1 {
		fmt.Println("ERR DEL <key> - Syntax error")
		return
	}

	key := args[0]
	msg := db.Delete(key)
	fmt.Println(msg)
}

func begin(args []string, db *database.Database) {
	if len(args) != 0 {
		fmt.Println("ERR This command do not receive arguments")
		return
	}

	msg := db.BeginTransaction()
	fmt.Println(msg)
}

func rollback(args []string, db *database.Database) {
	if len(args) != 0 {
		fmt.Println("ERR This command do not receive arguments")
		return
	}

	msg := db.Rollback()
	fmt.Println(msg)
}

func commit(args []string, db *database.Database) {
	if len(args) != 0 {
		fmt.Println("ERR This command do not receive arguments")
		return
	}

	msg := db.Commit()
	fmt.Println(msg)
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
