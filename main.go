package main

import (
	"cli-database/server"
	"log"
)

func main() {
	server := server.NewServer(":3000")
	log.Fatal(server.Start())

	// scanner := bufio.NewScanner(os.Stdin)
	// db := database.New()

	// for {
	// 	fmt.Print(">  ")
	// 	scanned := scanner.Scan()

	// 	if !scanned {
	// 		return
	// 	}

	// 	if scanner.Text() == ":q" {
	// 		cmd.CleanUp(db)
	// 		continue
	// 	}

	// 	if scanner.Text() == "" {
	// 		continue
	// 	}

	// 	input := lexer.Tokenize(scanner.Text())

	// 	cmd.Execute(input, db)
	// }
}
