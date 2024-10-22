package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Model untuk item di keranjang
type CartItem struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    ProductID primitive.ObjectID `bson:"product_id"`
    Quantity  int                `bson:"quantity"`
}
