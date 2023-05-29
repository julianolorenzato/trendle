package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/julianolorenzato/choosely/infra/network"
	"github.com/julianolorenzato/choosely/infra/persistence"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	persistence.InitialisePostgres()

	redisClient := persistence.InitialiseRedis()
	defer redisClient.Close()

	network.StartServer()
}
