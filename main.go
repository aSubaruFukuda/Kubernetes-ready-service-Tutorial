package main

import (
	"context"
	"github.com/subaruf/Kubernetes-ready-service-Tutorial/handlers"
	"github.com/subaruf/Kubernetes-ready-service-Tutorial/version"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// How to try it: PORT=8000 go run main.go
func main() {
	log.Printf("Starting the service...\ncommit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}
	r := handlers.Router(version.BuildTime, version.Commit, version.Release)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	go func() { log.Fatal(srv.ListenAndServe()) }()
	log.Print("the service is ready to listen and service...")
	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Printf("GO SIGINT...")
	case syscall.SIGTERM:
		log.Printf("GO SIGTERM...")
	}
	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}
