package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/product-service/models"
	"github.com/product-service/services"
)

var DB *sql.DB

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	sqlStatement := `
		INSERT INTO products (name, description, price, image_url)
		VALUES ($1, $2, $3, $4) RETURNING id`
	err = DB.QueryRow(sqlStatement, product.ProductName, product.ProductDescription, product.ProductPrice, product.ProductImages[0]).Scan(&product.ID)
	if err != nil {
		http.Error(w, "Failed to insert product", http.StatusInternalServerError)
		return
	}

	services.PublishToRabbitMQ(product)
	services.CacheProductInRedis(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	product, err := services.GetProductFromRedis(id)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
		return
	}

	product = &models.Product{}
	sqlStatement := `SELECT id, name, description, price, image_url FROM products WHERE id = $1`
	err = DB.QueryRow(sqlStatement, id).Scan(&product.ID, &product.ProductName, &product.ProductDescription, &product.ProductPrice, &product.ProductImages)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	services.CacheProductInRedis(*product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
