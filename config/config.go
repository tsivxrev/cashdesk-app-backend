package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[CONFIG] Cannot load config")
	}
}

func MongoURI() string {
	return os.Getenv("MONGO_URI")
}

func MongoDatabase() string {
	return os.Getenv("MONGO_DATABASE")
}
