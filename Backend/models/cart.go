package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type CartItem struct {
    ProductID primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
    Quantity  int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

type Cart struct {
    ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    UserID primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
    Items  []CartItem         `json:"items,omitempty" bson:"items,omitempty"`
}