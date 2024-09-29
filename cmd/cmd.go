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
	case "GET":
	case "BEGIN":
	case "ROLLBACK":
	case "COMMIT":
	default:
		fmt.Print("Unknown command: %s", command)
	}
}
