package queues

import (
	"context"
	"encoding/json"
	"github.com/julianolorenzato/choosely/core/domain"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisQueueConsumer struct {
	client *redis.Client
	voteDB domain.VoteDB
}

//type CallbackFn func(pollId string) map[string]int32

func (rqc *RedisQueueConsumer) receiveNewVote(e json.Encoder) {
	pubsub := rqc.client.Subscribe(context.Background(), "new_votes")

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

		err = e.Encode(pollFreshResults)
		if err != nil {
			log.Fatal(err)
		}
	}
}
