package queues

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisQueueProducer struct {
	client *redis.Client
}

func NewRedisQueueProducer() *RedisQueueProducer {
	return &RedisQueueProducer{
		client: redisClient,
	}
}

func (rqp *RedisQueueProducer) NotifyNewVote(pollID string) {
	channel := fmt.Sprintf("new_votes_in_%s", pollID)
	rqp.client.Publish(context.Background(), channel, "new_vote")
}
