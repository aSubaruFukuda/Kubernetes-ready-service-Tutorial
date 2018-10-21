package main

import (
	"log"
	"net/http"
	"./handlers"
)

func main() {
	log.Print("Starting the service...")
	router := handlers.Router()
	log.Print("the service is ready to listen and service...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
