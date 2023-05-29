package persistence

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitialiseRedis() *redis.Client {
	rdC := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	pong, err := rdC.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Redis connection failed.", err)
	}

	log.Println("Redis sucessfully connected. Ping:", pong)

	return rdC
}