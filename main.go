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

		result := lexer.New(scanner.Text())

		fmt.Print(result)
		fmt.Print("\n")
		fmt.Print(len(result))
		fmt.Print("\n")
	}
}
