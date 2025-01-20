package cmd

import (
	"cli-database/database"
	"fmt"
)

func get(args []string, db *database.Database) (string, error) {
	if len(args) != 1 {
		err := fmt.Errorf("ERR GET <key> - Syntax error")
		return "", err
	}

	key := args[0]
	msg := db.Get(key)

	return msg, nil
}

func set(args []string, db *database.Database) (string, error) {
	if len(args) != 2 {
		err := fmt.Errorf("ERR SET <key> - <value> - Syntax error")
		return "", err
	}

	key := args[0]
	value := args[1]

	msg := db.Set(key, value)
	return msg, nil
}

func mget(keys []string, db *database.Database) (string, error) {
	if len(keys) == 0 {
		err := fmt.Errorf("ERR MGET need at least 1 keys - Syntax error")
		return "", err
	}

	values := db.Mget(keys)

	var result string

	for i, v := range values {
		result += fmt.Sprintf("key: %s, value: %s \n", keys[i], v)
	}
	return result, nil
}

func mset(args []string, db *database.Database) (string, error) {
	if len(args)%2 != 0 {
		err := fmt.Errorf("ERR args of mset cannot be odd - Synax error")
		return "", err
	}

	msg := db.Mset(args)
	return msg, nil
}

func incr(args []string, db *database.Database) (string, error) {
	if len(args) != 1 {
		err := fmt.Errorf("ERR INCR <key> - Syntax error")
		return "", err
	}

	key := args[0]
	msg := db.Incr(key)
	return msg, nil
}

func decr(args []string, db *database.Database) (string, error) {
	if len(args) != 1 {
		err := fmt.Errorf("ERR DECR <key> - Syntax error")
		return "", err
	}

	key := args[0]
	msg := db.Decr(key)
	return msg, nil
}
