package main

import (
	"github.com/subaruf/Kubernetes-ready-service-Tutorial/handlers"
	"github.com/subaruf/Kubernetes-ready-service-Tutorial/version"
	"log"
	"net/http"
	"os"
)

// How to try it: PORT=8000 go run main.go
func main() {
	log.Printf("Starting the service...\ncommit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}
	r := handlers.Router(version.BuildTime, version.Commit, version.Release)
	log.Print("the service is ready to listen and service...")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
