package cmd

import (
	"cli-database/database"
	"fmt"
	"strings"
)

func Execute(input []string, db *database.Database) (string, error) {
	command := strings.ToUpper(input[0])
	args := input[1:]

	switch command {
	case "GET":
		return get(args, db)
	case "MGET":
		return mget(args, db)
	case "SET":
		return set(args, db)
	case "MSET":
		return mset(args, db)
	case "BEGIN":
		return begin(args, db)
	case "ROLLBACK":
		return rollback(args, db)
	case "COMMIT":
		return commit(args, db)
	case "DEL":
		return delete(args, db)
	case "COPY":
		return copy(args, db)
	case "INCR":
		return incr(args, db)
	case "DECR":
		return decr(args, db)
	case "TTL":
		return ttl(args, db)
	case "PERSIST":
		return persist(args, db)
	case "EXPIRE":
		return expire(args, db)
	case "RENAME":
		return rename(args, db)
	case "EXISTS":
		return exists(args, db)
	default:
		err := fmt.Errorf("unknow command %s \n", command)
		return "", err
	}
}

func get(args []string, db *database.Database) (string, error) {
	if len(args) != 1 {
		err := fmt.Errorf("ERR GET <key> - Syntax error")
		return "", err
	}

	key := args[0]
	msg := db.Get(key)

	return msg, nil
}

func mget(keys []string, db *database.Database) (string, error) {
	if len(keys) == 0 {
		err := fmt.Errorf("ERR MGET need at least 1 keys - Syntax error")
		return "", err
	}

	// values := db.Mget(keys)

	// for i, v := range values {
	// 	fmt.Printf("key: %s, value: %s \n", keys[i], v)
	// }
	return "empty string for now", nil
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

func mset(args []string, db *database.Database) (string, error) {
	if len(args)%2 != 0 {
		err := fmt.Errorf("ERR args of mset cannot be odd - Synax error")
		return "", err
	}

	msg := db.Mset(args)
	return msg, nil
}

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
