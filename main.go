package main

import (
	"log"
	"net/http"

	"github.com/Fr3shDev/freshserve/handlers"
	"github.com/Fr3shDev/freshserve/middleware"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Register middleware - here, we add our logging middleware
	router.Use(middleware.LoggingMiddleware)

	// Define routes and associate them with handler functions
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")

	// Define the port to listen on
	port := ":8080"
	log.Printf("Server is starting on port %s", port)

	// Start the HTTP server
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}