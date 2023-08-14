package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// Load environment vars
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func Env(key string) string {
	return os.Getenv(key)
}
