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

		input := lexer.New(scanner.Text())

		err := cmd.Execute(input, db)

		if err != nil {
			fmt.Print(err)
			fmt.Print("\n")
		}
	}
}
