package main

import (
	"domain-to-ip/app"
	"log"
	"os"
)

func main() {
	application := app.Generate()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
