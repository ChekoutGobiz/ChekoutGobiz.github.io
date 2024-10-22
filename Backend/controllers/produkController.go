package controllers

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "yourapp/models"
)

type ProductController struct {
    ProductCollection *mongo.Collection
}

func NewProductController(productCollection *mongo.Collection) *ProductController {
    return &ProductController{ProductCollection: productCollection}
}

// Buat produk baru
func (p *ProductController) CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := p.ProductCollection.InsertOne(ctx, product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan produk"})
        return
    }

    c.JSON(http.StatusCreated, product)
}

// Ambil semua produk
func (p *ProductController) GetAllProducts(c *gin.Context) {
    var products []models.Product
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := p.ProductCollection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan produk"})
        return
    }

    if err := cursor.All(ctx, &products); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses data produk"})
        return
    }

    c.JSON(http.StatusOK, products)
}
