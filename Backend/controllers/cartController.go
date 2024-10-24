package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/kuyjajan/kuyjajan-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"github.com/joho/godotenv"
)

var cartCollection *mongo.Collection
var cartClient *mongo.Client

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Ambil MONGODB_URI dari environment
	mongoURI := os.Getenv("MONGODB_URI")

	// Opsi koneksi MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)
	cartClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Cek koneksi MongoDB
	err = cartClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Inisialisasi koleksi cart
	cartCollection = cartClient.Database("jajankuy").Collection("carts")
}

// AddToCart menambahkan item ke dalam keranjang
func AddToCart(w http.ResponseWriter, r *http.Request) {
	var cartItem models.CartItem
	var cart models.Cart

	if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Assume user is authenticated and we get the user ID (This should be handled via JWT authentication)
	userID, _ := primitive.ObjectIDFromHex("USER_ID") // Gantilah dengan ID pengguna dari JWT

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Temukan keranjang user
	err := cartCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&cart)
	if err == mongo.ErrNoDocuments {
		// Jika keranjang belum ada, buat baru
		cart.ID = primitive.NewObjectID()
		cart.UserID = userID
		cart.Items = []models.CartItem{cartItem}
		_, err := cartCollection.InsertOne(ctx, cart)
		if err != nil {
			http.Error(w, "Failed to add to cart", http.StatusInternalServerError)
			return
		}
	} else {
		// Tambahkan item ke keranjang yang sudah ada
		cart.Items = append(cart.Items, cartItem)
		_, err := cartCollection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$set": bson.M{"items": cart.Items}})
		if err != nil {
			http.Error(w, "Failed to update cart", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

// GetCart mengambil keranjang user
func GetCart(w http.ResponseWriter, r *http.Request) {
	// Assume user is authenticated and we get the user ID (handled via JWT)
	userID, _ := primitive.ObjectIDFromHex("USER_ID") // Gantilah dengan ID pengguna dari JWT

	var cart models.Cart
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := cartCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&cart)
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}