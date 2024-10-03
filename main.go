package main

import (
	"bufio"
	"cli-database/cmd"
	"cli-database/database"
	"cli-database/lexer"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	db := database.New()

	for {
		fmt.Print(">  ")
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		if scanner.Text() == ":q" {
			cmd.CleanUp(db)
			continue
		}

		if scanner.Text() == "" {
			continue
		}

		input := lexer.Tokenize(scanner.Text())

		cmd.Execute(input, db)
	}
}
