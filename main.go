package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/product-service/config"
	"github.com/product-service/handlers"
	"github.com/product-service/services"
	_ "github.com/lib/pq"
)

func main() {
	config.LoadEnv()

	// Initialize PostgreSQL
	var err error
	handlers.DB, err = sql.Open("postgres", "user=your_db_user password=your_db_password dbname=your_db_name sslmode=require")
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	defer handlers.DB.Close()

	// Initialize RabbitMQ and Redis
	services.InitRabbitMQ()
	defer services.RabbitMQConn.Close()
	defer services.RabbitMQChannel.Close()

	services.InitRedis()
	defer services.RedisClient.Close()

	// Router setup
	r := mux.NewRouter()
	r.HandleFunc("/products", handlers.CreateProductHandler).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.GetProductByIDHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
