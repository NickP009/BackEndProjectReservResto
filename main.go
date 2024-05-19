package main

import (
	"log"
	"net/http"

	"backend/database"
	"backend/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize the database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize router
	r := mux.NewRouter()

	// Register handlers
	r.HandleFunc("/register", handlers.Register(db)).Methods("POST")
	r.HandleFunc("/login", handlers.Login(db)).Methods("POST")
	r.HandleFunc("/reserve", handlers.MakeReservation(db)).Methods("POST")

	// Enable CORS
	handler := cors.Default().Handler(r)

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
