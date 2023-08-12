package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/julianolorenzato/choosely/network"
)

func main() {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Start app server
	port := os.Getenv("PORT")
	network.NewHTTPServer(":" + port).Start()
}
