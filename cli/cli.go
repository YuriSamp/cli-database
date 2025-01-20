package cli

import (
	"bufio"
	"cli-database/cmd"
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

		if !scanned || scanner.Text() == "" {
			return
		}

		input := lexer.Tokenize(scanner.Text())

		msg, err := cmd.Execute(input, db)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(msg)
		}
	}
}
