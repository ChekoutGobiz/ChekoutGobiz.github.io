package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Model untuk produk
type Product struct {
    ID            primitive.ObjectID `bson:"_id,omitempty"`
    Name          string             `bson:"name"`
    Description   string             `bson:"description"`
    Rating        float64            `bson:"rating"`
    DiscountedPrice float64          `bson:"discounted_price"`
    OriginalPrice  float64           `bson:"original_price"`
}
