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
	// Endpoint untuk region
	r.HandleFunc("/regions", controllers.CreateRegion).Methods("POST")
	r.HandleFunc("/regions", controllers.GetRegions).Methods("GET")

	// Daftarkan rute untuk produk
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products/region", controllers.GetProductsByRegion).Methods("GET") // Rute baru untuk pencarian produk berdasarkan region
	// Rute untuk cart
	r.HandleFunc("/cart", controllers.AddToCart).Methods("POST")
	r.HandleFunc("/cart", controllers.GetCart).Methods("GET")
	r.HandleFunc("/cart/item", controllers.UpdateCartItem).Methods("PUT")
	r.HandleFunc("/cart/item", controllers.RemoveCartItem).Methods("DELETE")

	// Configure CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://127.0.0.1:5501"}),        // Allow Live Server origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), // Allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
	// Jalankan server
	log.Println("Server berjalan di port 8080")
}
