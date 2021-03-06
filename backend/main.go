package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	// Get the "PORT" env variable
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := store.NewRouter() // create routes

	// These two lines are important if you're designing a front-end to utilise this API methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))

	// if everything works fine, proceed with transaction in NAVER PAY
	// something to pop up NAVER PAY html file
	// returns transaction result and sends it to database
}
