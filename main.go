package main

import (
	"cli-database/cli"
	"cli-database/database"
	"cli-database/server"
	"log"
	"os"
)

func main() {
	args := os.Args
	db := database.New()

	if len(args) > 1 && args[1] == "--tcp" {
		server := server.NewServer(":3000", db)
		log.Fatal(server.Start())
		return
	} else {
		// for simplicity i don't care about a lot of things right now
		cli.StartCli(db)
	}
}
