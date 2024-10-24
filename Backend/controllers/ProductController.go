package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/kuyjajan/kuyjajan-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection

// Initialize collection in init function
func init() {
	productCollection = client.Database("jajankuy").Collection("products")
}

// CreateProduct handles the creation of a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if !product.Validate() {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	product.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		http.Error(w, "Failed to insert product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetProducts retrieves all products from the database
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product models.Product
		cursor.Decode(&product)
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}