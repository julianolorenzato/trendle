package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/julianolorenzato/choosely/infra/network"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}



	network.StartServer()
}
