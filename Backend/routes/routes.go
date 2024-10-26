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
	router.HandleFunc("/products/region", controllers.GetProductsByRegion).Methods("GET") // Rute baru untuk pencarian produk berdasarkan region

	// Endpoint untuk region
	router.HandleFunc("/regions", controllers.CreateRegion).Methods("POST")
	router.HandleFunc("/regions", controllers.GetRegions).Methods("GET")

	// Rute untuk cart
	router.HandleFunc("/cart", controllers.AddToCart).Methods("POST")
	router.HandleFunc("/cart", controllers.GetCart).Methods("GET")
	router.HandleFunc("/cart/item", controllers.UpdateCartItem).Methods("PUT")
	router.HandleFunc("/cart/item", controllers.RemoveCartItem).Methods("DELETE")

	return router
}
