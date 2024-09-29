package main

import (
	"bufio"
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

		lexer.New(scanner.Text())
	}
}
