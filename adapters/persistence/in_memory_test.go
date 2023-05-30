package persistence_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/julianolorenzato/choosely/domain/poll"
	"github.com/julianolorenzato/choosely/adapters/persistence"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryPollRepository(t *testing.T) {
	t.Run("It should get a poll from memory", func(t *testing.T) {
		ids := []string{uuid.NewString(), uuid.NewString(), uuid.NewString()}

		iMPR := &persistence.InMemoryPollRepository{
			Polls: []*poll.Poll{
				{ID: ids[0]},
				{ID: ids[1]},
				{ID: ids[2]},
			},
		}

		p, err := iMPR.GetByID(ids[1])

		assert.Nil(t, err)
		assert.NotNil(t, p)
		assert.Equal(t, p, iMPR.Polls[1])
	})

	t.Run("It should return nil if the poll does not exists in memory", func(t *testing.T) {
		iMPR := &persistence.InMemoryPollRepository{
			Polls: []*poll.Poll{},
		}

		p, err := iMPR.GetByID(uuid.NewString())

		assert.Nil(t, err)
		assert.Nil(t, p)
	})

	t.Run("It should create a new poll in the memory", func(t *testing.T) {
		iMPR := &persistence.InMemoryPollRepository{
			Polls: []*poll.Poll{},
		}

		poll := &poll.Poll{ID: uuid.NewString()}

		err := iMPR.Create(poll)

		assert.Nil(t, err)
		assert.Contains(t, iMPR.Polls, poll)
	})

	t.Run("It should create a new poll in the memory if doest not exists", func(t *testing.T) {
		iMPR := &persistence.InMemoryPollRepository{
			Polls: []*poll.Poll{},
		}
		
		poll := &poll.Poll{ID: uuid.NewString()}
		
		err := iMPR.Save(poll)
		
		assert.Nil(t, err)
		assert.Contains(t, iMPR.Polls, poll)
	})
	
	t.Run("It should update a poll that already exists in memory", func(t *testing.T) {
		p := &poll.Poll{ID: uuid.NewString(), Question: "Before"}

		iMPR := &persistence.InMemoryPollRepository{
			Polls: []*poll.Poll{p},
		}

		p.Question = "After"

		err := iMPR.Save(p)

		assert.Nil(t, err)
		assert.Len(t, iMPR.Polls, 1)
		assert.Equal(t, iMPR.Polls[0].Question, "After")
	})
}
