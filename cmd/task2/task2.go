package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ronen25/task2/service"
)

func main() {
	// Get server port
	serverPort, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		log.Fatalln("Error: Environment variable `SERVER_PORT` is missing")
	}

	service := service.NewService()

	// Initialize router
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/", service.QueryPrintingHandler)

	bindAddress := fmt.Sprintf(":%s", serverPort)
	log.Printf("Listening on address %s", bindAddress)
	if err := http.ListenAndServe(bindAddress, router); err != nil {
		log.Fatalf("Error setting up HTTP server: %v", err)
	}
}
