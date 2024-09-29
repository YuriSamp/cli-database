package main

import (
	"bufio"
	"cli-database/cmd"
	"cli-database/lexer"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">  ")
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		input := lexer.New(scanner.Text())

		cmd.Execute(input)
	}
}
