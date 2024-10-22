package routes

import (
	"github.com/gorilla/mux"
	"github.com/kuyjajan/kuyjajan-backend/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterRoutes mengatur semua rute untuk aplikasi
func RegisterRoutes(productCollection, cartCollection *mongo.Collection) *mux.Router {
	router := mux.NewRouter()

	// Routes untuk register dan login
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	// Routes untuk produk
	productController := controllers.NewProductController(productCollection)
	router.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	router.HandleFunc("/products", productController.GetAllProducts).Methods("GET")

	// Routes untuk keranjang
	cartController := controllers.NewCartController(cartCollection)
	router.HandleFunc("/cart", cartController.AddToCart).Methods("POST")
	router.HandleFunc("/cart", cartController.GetCartItems).Methods("GET")

	return router
}
