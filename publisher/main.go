package main

import (
	"log"

	"publisher/services"
)

func main() {
	handlers, err := services.NewHandlers()
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := handlers.Publisher(); err != nil {
		log.Fatal(err)
		return
	}
}
