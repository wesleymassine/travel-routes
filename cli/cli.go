package main

import (
	"log"
	"os"
	"travel-routes/cli/app"
)

func main() {
	application := app.Cli()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
