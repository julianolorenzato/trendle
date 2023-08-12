package persistence_test

import (
	"github.com/julianolorenzato/choosely/core/domain"
	"github.com/julianolorenzato/choosely/persistence"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	a := assert.New(t)

	t.Run("It should return true for a poll that exists", func(t *testing.T) {
		ids := []string{uuid.NewString(), uuid.NewString(), uuid.NewString()}

		iMPR := &persistence.InMemoryPollDB{
			Polls: []*domain.Poll{
				{ID: ids[0]},
				{ID: ids[1]},
				{ID: ids[2]},
			},
		}

		exists := iMPR.Exists(ids[0])

		a.True(exists)
	})

	t.Run("It should return false for a poll that not exists", func(t *testing.T) {
		ids := []string{uuid.NewString(), uuid.NewString(), uuid.NewString()}

		iMPR := &persistence.InMemoryPollDB{
			Polls: []*domain.Poll{
				{ID: ids[0]},
				{ID: ids[1]},
				{ID: ids[2]},
			},
		}

		exists := iMPR.Exists("not an uuid")

		a.False(exists)
	})
}

func TestGetByID(t *testing.T) {
	a := assert.New(t)

	t.Run("It should get a poll from memory", func(t *testing.T) {
		ids := []string{uuid.NewString(), uuid.NewString(), uuid.NewString()}

		iMPR := &persistence.InMemoryPollDB{
			Polls: []*domain.Poll{
				{ID: ids[0]},
				{ID: ids[1]},
				{ID: ids[2]},
			},
		}

		p, err := iMPR.GetByID(ids[1])

		a.Nil(err)
		a.NotNil(p)
		a.Equal(p, iMPR.Polls[1])
	})

	t.Run("It should return nil if the poll does not exists in memory", func(t *testing.T) {
		iMPR := &persistence.InMemoryPollDB{
			Polls: []*domain.Poll{},
		}

		p, err := iMPR.GetByID(uuid.NewString())

		a.Nil(err)
		a.Nil(p)
	})

}

func TestCreate(t *testing.T) {
	a := assert.New(t)

	t.Run("It should create a new poll in the memory", func(t *testing.T) {
		iMPR := &persistence.InMemoryPollDB{
			Polls: []*domain.Poll{},
		}

		poll := &domain.Poll{ID: uuid.NewString()}

		err := iMPR.Create(poll)

		a.Nil(err)
		a.Contains(iMPR.Polls, poll)
	})
}

func TestSave(t *testing.T) {
	a := assert.New(t)

	t.Run("It should create a new poll in the memory if doest not exists", func(t *testing.T) {
		iMPR := &persistence.InMemoryPollDB{
			Polls: []*domain.Poll{},
		}

		poll := &domain.Poll{ID: uuid.NewString()}

		err := iMPR.Save(poll)

		a.Nil(err)
		a.Contains(iMPR.Polls, poll)
	})

	t.Run("It should update a poll that already exists in memory", func(t *testing.T) {
		p := &domain.Poll{ID: uuid.NewString(), Question: "Before"}

		iMPR := &persistence.InMemoryPollDB{
			Polls: []*domain.Poll{p},
		}

		p.Question = "After"

		err := iMPR.Save(p)

		a.Nil(err)
		a.Len(iMPR.Polls, 1)
		a.Equal(iMPR.Polls[0].Question, "After")
	})
}
