package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	Stock       int                `json:"stock,omitempty" bson:"stock,omitempty"`
}

func (product *Product) Validate() bool {
	// Implementasi validasi produk jika diperlukan, seperti memastikan nama, harga, dan stok diisi
	return product.Name != "" && product.Price > 0 && product.Stock >= 0
}