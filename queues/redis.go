package queues

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var redisClient *redis.Client

func init() {
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

	redisClient = rdC
}
