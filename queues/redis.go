package queues

import (
	"context"
	"github.com/julianolorenzato/choosely/config"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisClient *redis.Client

func init() {
	rdC := redis.NewClient(&redis.Options{
		Addr:     config.Env("REDIS_ADDR"),
		Password: config.Env("REDIS_PASS"),
		DB:       0,
	})

	pong, err := rdC.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Redis connection failed. ", err)
	}

	log.Println("Redis successfully connected. Ping:", pong)

	redisClient = rdC
}
