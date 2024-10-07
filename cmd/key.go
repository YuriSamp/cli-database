package cmd

import (
	"cli-database/database"
	"fmt"
)

func ttl(args []string, db *database.Database) {
	if len(args) != 1 {
		fmt.Println("ERR TTL <key> - Syntax error")
		return
	}

	key := args[0]
	msg := db.TTL(key)
	fmt.Println(msg)
}

func persist(args []string, db *database.Database) {
	if len(args) != 1 {
		fmt.Println("ERR PERSIST <key> - Syntax error")
		return
	}

	key := args[0]
	msg := db.Persist(key)
	fmt.Println(msg)
}

func expire(args []string, db *database.Database) {
	if len(args) != 2 {
		fmt.Println("ERR EXPIRE <KEY> <TIME> - Syntax error")
		return
	}

	key := args[0]
	time := args[1]

	msg := db.Expire(key, time)
	fmt.Println(msg)
}

func copy(args []string, db *database.Database) {
	if len(args) != 2 {
		fmt.Println("ERR COPY <source> <destination> - Syntax error")
		return
	}

	source := args[0]
	destination := args[1]

	msg := db.Copy(source, destination)
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

func rename(args []string, db *database.Database) {
	if len(args) != 2 {
		fmt.Println("ERR RENAME <source> <destination> - Syntax error")
		return
	}

	source := args[0]
	destionation := args[1]

	msg := db.Rename(source, destionation)
	fmt.Println(msg)
}

func exists(args []string, db *database.Database) {
	if len(args) == 0 {
		fmt.Println("ERR EXISTIS key [key ...] - Syntax error")
		return
	}

	msg := db.Exists(args)
	fmt.Println(msg)
}
