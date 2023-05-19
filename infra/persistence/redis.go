package persistence

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDR"),
	Password: "",
	DB:       0,
})

var ctx = context.Background()

func TestConnection() {
	_, err := redisClient.Ping(ctx).Result()

	if err != nil {
		log.Println(err)
	}
}

func Set(str any) {
}