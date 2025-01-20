package cmd

import (
	"bufio"
	"cli-database/database"
	"cli-database/lexer"
	"fmt"
	"os"
)

func StartCli(db *database.Database) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">  ")
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		if scanner.Text() == ":q" {
			CleanUp(db)
			continue
		}

		if scanner.Text() == "" {
			continue
		}

		input := lexer.Tokenize(scanner.Text())

		Execute(input, db)
	}
}
