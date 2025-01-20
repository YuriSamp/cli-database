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
		err := fmt.Errorf("unknow command %s", command)
		return "", err
	}
}
