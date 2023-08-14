package queues

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julianolorenzato/choosely/core/domain"
	"github.com/redis/go-redis/v9"
	"io"
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

func (rqc *RedisQueueConsumer) SubscribeToPollChannel(pollID string, w io.WriteCloser) {
	channel := fmt.Sprintf("new_votes_in_%s", pollID)
	pubsub := rqc.client.Subscribe(context.Background(), channel)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		pollId := msg.Payload

		pollFreshResults, err := rqc.voteDB.GetResults(pollId)
		if err != nil {
			log.Fatal(err)
		}

		err = json.NewEncoder(w).Encode(pollFreshResults)
		if err != nil {
			log.Fatal(err)
		}
		// Need close to websocket flushes the message
		w.Close()
	}
}
