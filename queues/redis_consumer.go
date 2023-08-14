package queues

import (
	"context"
	"fmt"
	"github.com/julianolorenzato/choosely/core/domain"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisQueueConsumer struct {
	client *redis.Client
	voteDB domain.VoteDB
}

func NewRedisQueueConsumer() *RedisQueueConsumer {
	return &RedisQueueConsumer{
		client: redisClient,
	}
}

func (rqc *RedisQueueConsumer) SubscribeToPollChannel(pollID string, callback func()) {
	channel := fmt.Sprintf("new_votes_in_%s", pollID)
	pubsub := rqc.client.Subscribe(context.Background(), channel)
	defer pubsub.Close()
	for {
		_, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		callback()
	}
}
