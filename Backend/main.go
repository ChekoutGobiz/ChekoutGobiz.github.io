package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/kuyjajan/kuyjajan-backend/controllers"
)

func main() {
	// Koneksi ke MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	productCollection := client.Database("yourdb").Collection("products")
	cartCollection := client.Database("yourdb").Collection("cart")

	// Inisialisasi router
	r := mux.NewRouter()

	// Routes untuk register dan login
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Routes untuk produk
	productController := controllers.NewProductController(productCollection)
	r.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	r.HandleFunc("/products", productController.GetAllProducts).Methods("GET")

	// Routes untuk keranjang
	cartController := controllers.NewCartController(cartCollection)
	r.HandleFunc("/cart", cartController.AddToCart).Methods("POST")
	r.HandleFunc("/cart", cartController.GetCartItems).Methods("GET")

	// Configure CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://127.0.0.1:5501"}),        // Allow Live Server origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), // Allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
