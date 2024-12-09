package config

import (
	"log"
	"os"
)

func LoadEnv() {
	os.Setenv("DB_USER", "your_db_user")
	os.Setenv("DB_PASSWORD", "your_db_password")
	os.Setenv("DB_NAME", "your_db_name")
	os.Setenv("REDIS_HOST", "localhost:6379")
	os.Setenv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	log.Println("Environment variables loaded")
}
