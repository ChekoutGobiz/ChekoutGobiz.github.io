// routes.go
package routes

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kuyjajan/kuyjajan-backend/controllers"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes() http.Handler {
	r := mux.NewRouter()

	// Authentication routes
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Region routes
	r.HandleFunc("/regions", controllers.CreateRegion).Methods("POST")
	r.HandleFunc("/regions", controllers.GetRegions).Methods("GET")

	// Product routes
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products/region", controllers.GetProductsByRegion).Methods("GET")

	// Cart routes (API for handling the cart and checkout/receipt generation)
	r.HandleFunc("/cart", controllers.AddToCart).Methods("POST")        // Add items to the cart
	r.HandleFunc("/cart", controllers.GetCart).Methods("GET")           // Retrieve the cart content (acts as a receipt)
	r.HandleFunc("/cart/item", controllers.UpdateCartItem).Methods("PUT") // Update a specific cart item
	r.HandleFunc("/cart/item", controllers.RemoveCartItem).Methods("DELETE") // Remove a specific item from the cart

	// Configure CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://127.0.0.1:5501"}),              // Allow Live Server origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),       // Allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),      // Allowed headers for CORS
	)

	return corsHandler(r)
}
