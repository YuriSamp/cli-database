package cmd

import (
	"strings"
)

func Execute(input []string) {
	command := strings.ToUpper(input[0])
	args := input[1:]

	switch command {
	case "SET":
	case "GET":
	case "BEGIN":
	case "ROLLBACK":
	case "COMMIT":
	}
}
