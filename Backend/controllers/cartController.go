package controllers

import (
    "context"
    "net/http"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "yourapp/models" // Ganti dengan path yang sesuai
)

type CartController struct {
    CartCollection *mongo.Collection
}

func NewCartController(cartCollection *mongo.Collection) *CartController {
    return &CartController{CartCollection: cartCollection}
}

// Tambahkan item ke keranjang
func (cc *CartController) AddToCart(w http.ResponseWriter, r *http.Request) {
    var cartItem models.CartItem
    if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := cc.CartCollection.InsertOne(ctx, cartItem)
    if err != nil {
        http.Error(w, "Gagal menambahkan ke keranjang", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(cartItem)
}

// Ambil semua item di keranjang
func (cc *CartController) GetCartItems(w http.ResponseWriter, r *http.Request) {
    var cartItems []models.CartItem
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := cc.CartCollection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, "Gagal mendapatkan item keranjang", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx) // Tutup cursor setelah selesai

    if err := cursor.All(ctx, &cartItems); err != nil {
        http.Error(w, "Gagal memproses data keranjang", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(cartItems)
}
