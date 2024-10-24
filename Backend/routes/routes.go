package routes

import (
	"github.com/gorilla/mux"
	"github.com/kuyjajan/kuyjajan-backend/controllers"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	// Product routes
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")

	// Cart routes
	router.HandleFunc("/cart", controllers.AddToCart).Methods("POST")
	router.HandleFunc("/cart", controllers.GetCart).Methods("GET")

	return router
}