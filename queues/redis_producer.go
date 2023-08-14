package queues

import (
	"context"
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

func (rqp *RedisQueueProducer) notifyNewVote(pollID string) error {
	rqp.client.Publish(context.Background(), "new_votes", pollID)
	return nil
}
