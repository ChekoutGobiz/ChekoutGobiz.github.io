package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kuyjajan/kuyjajan-backend/controllers"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Configure CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://127.0.0.1:5501"}),        // Allow Live Server origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), // Allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
