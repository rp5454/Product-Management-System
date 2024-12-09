package services

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/product-service/models"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}

func CacheProductInRedis(product models.Product) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	productJSON, _ := json.Marshal(product)
	err := RedisClient.Set(ctx, strconv.Itoa(product.ID), productJSON, 10*time.Minute).Err()
	if err != nil {
		log.Println("Failed to cache product in Redis:", err)
	}
}

func GetProductFromRedis(id string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	productJSON, err := RedisClient.Get(ctx, id).Result()
	if err != nil {
		return nil, err
	}

	var product models.Product
	err = json.Unmarshal([]byte(productJSON), &product)
	return &product, err
}
