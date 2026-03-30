package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"reader/services"
)

func main() {
	handlers, err := services.NewHandlers()
	if err != nil {
		log.Fatal(err)
		return
	}

	sub, err := handlers.Reader()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer handlers.Pg.CloseConnect()
	defer sub.Unsubscribe()

	log.Println("reader subscribed and waiting for messages")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("shutting down reader")
}
