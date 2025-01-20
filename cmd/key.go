package cmd

import (
	"cli-database/database"
	"fmt"
)

func ttl(args []string, db *database.Database) (string, error) {
	if len(args) != 1 {
		err := fmt.Errorf("ERR TTL <key> - Syntax error")
		return "", err
	}

	key := args[0]
	msg := db.TTL(key)
	return msg, nil
}

func persist(args []string, db *database.Database) (string, error) {
	if len(args) != 1 {
		err := fmt.Errorf("ERR PERSIST <key> - Syntax error")
		return "", err
	}

	key := args[0]
	msg := db.Persist(key)
	return msg, nil
}

func expire(args []string, db *database.Database) (string, error) {
	if len(args) != 2 {
		err := fmt.Errorf("ERR EXPIRE <KEY> <TIME> - Syntax error")
		return "", err
	}

	key := args[0]
	time := args[1]

	msg := db.Expire(key, time)
	return msg, nil
}

func copy(args []string, db *database.Database) (string, error) {
	if len(args) != 2 {
		err := fmt.Errorf("ERR COPY <source> <destination> - Syntax error")
		return "", err
	}

	source := args[0]
	destination := args[1]

	msg := db.Copy(source, destination)
	return msg, nil
}

func delete(args []string, db *database.Database) (string, error) {
	if len(args) != 1 {
		err := fmt.Errorf("ERR DEL <key> - Syntax error")
		return "", err
	}

	key := args[0]
	msg := db.Delete(key)
	return msg, nil
}

func rename(args []string, db *database.Database) (string, error) {
	if len(args) != 2 {
		err := fmt.Errorf("ERR RENAME <source> <destination> - Syntax error")
		return "", err
	}

	source := args[0]
	destionation := args[1]

	msg := db.Rename(source, destionation)
	return msg, nil
}

func exists(args []string, db *database.Database) (string, error) {
	if len(args) == 0 {
		err := fmt.Errorf("ERR EXISTIS key [key ...] - Syntax error")
		return "", err
	}

	msg := db.Exists(args)
	fmt.Println(msg)
	return msg, nil
}
