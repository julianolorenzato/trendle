package persistence

import (
	"github.com/julianolorenzato/choosely/core/domain"
)

type InMemoryPollDB struct {
	Polls []*domain.Poll
}

func NewInMemoryPollDB() *InMemoryPollDB {
	return &InMemoryPollDB{
		Polls: make([]*domain.Poll, 0),
	}
}

func (r *InMemoryPollDB) Exists(ID string) bool {
	for i := range r.Polls {
		if r.Polls[i].ID == ID {
			return true
		}
	}

	return false
}

func (r *InMemoryPollDB) GetByID(ID string) (*domain.Poll, error) {
	for i := range r.Polls {
		if r.Polls[i].ID == ID {
			return r.Polls[i], nil
		}
	}

	return nil, nil
}

func (r *InMemoryPollDB) Create(poll *domain.Poll) error {
	r.Polls = append(r.Polls, poll)
	return nil
}

func (r *InMemoryPollDB) Save(poll *domain.Poll) error {
	for i := range r.Polls {
		if r.Polls[i].ID == poll.ID {
			r.Polls[i] = poll
			return nil
		}
	}

	r.Polls = append(r.Polls, poll)

	return nil
}

type InMemoryVoteDB struct {
	Votes []*domain.Vote
}

func NewInMemoryVoteDB() *InMemoryVoteDB {
	return &InMemoryVoteDB{
		Votes: make([]*domain.Vote, 0),
	}
}

func (r *InMemoryVoteDB) Create(vote *domain.Vote) error {
	r.Votes = append(r.Votes, vote)
	return nil
}

func (r *InMemoryVoteDB) GetByID(ID string) (*domain.Vote, error) {
	for i := range r.Votes {
		if r.Votes[i].ID == ID {
			return r.Votes[i], nil
		}
	}

	return nil, nil
}

func (r *InMemoryVoteDB) GetResults(pollID string) (map[string]uint, error) {
	res := make(map[string]uint)

	for i := range r.Votes {
		if r.Votes[i].PollID == pollID {

			for _, option := range r.Votes[i].ChoosenOptions {
				res[option]++
			}
		}
	}

	return res, nil
}
