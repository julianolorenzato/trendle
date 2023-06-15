package poll_test

import (
	"testing"
	"time"

	"github.com/julianolorenzato/choosely/domain/poll"
	"github.com/julianolorenzato/choosely/adapters/persistence"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewPoll(t *testing.T) {
	a := assert.New(t)

	t.Run("Should create a new poll", func(t *testing.T) {
		s := &poll.PollService{
			PollRepo: persistence.NewInMemoryPollRepository(),
			VoteRepo: persistence.NewInMemoryVoteRepository(),
		}

		err := s.CreateNewPoll(poll.CreateNewPollDTO{
			Question: "Some question",
			Options: []string{
				"opt0", "opt1", "opt2",
			},
			NumberOfChoices: 2,
			IsPermanent: true,
			ExpiresAt: time.Now(),
		})

		a.Nil(err)
		
	})
}
