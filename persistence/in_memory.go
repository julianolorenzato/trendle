package persistence

import (
	"github.com/julianolorenzato/choosely/core/domain"
)

type InMemoryPollRepository struct {
	Polls []*domain.Poll
}

func NewInMemoryPollRepository() *InMemoryPollRepository {
	return &InMemoryPollRepository{
		Polls: make([]*domain.Poll, 0),
	}
}

func (r *InMemoryPollRepository) Exists(ID string) bool {
	for i := range r.Polls {
		if r.Polls[i].ID == ID {
			return true
		}
	}

	return false
}

func (r *InMemoryPollRepository) GetByID(ID string) (*domain.Poll, error) {
	for i := range r.Polls {
		if r.Polls[i].ID == ID {
			return r.Polls[i], nil
		}
	}

	return nil, nil
}

func (r *InMemoryPollRepository) Create(poll *domain.Poll) error {
	r.Polls = append(r.Polls, poll)
	return nil
}

func (r *InMemoryPollRepository) Save(poll *domain.Poll) error {
	for i := range r.Polls {
		if r.Polls[i].ID == poll.ID {
			r.Polls[i] = poll
			return nil
		}
	}

	r.Polls = append(r.Polls, poll)

	return nil
}

type InMemoryVoteRepository struct {
	Votes []*domain.Vote
}

func NewInMemoryVoteRepository() *InMemoryVoteRepository {
	return &InMemoryVoteRepository{
		Votes: make([]*domain.Vote, 0),
	}
}

func (r *InMemoryVoteRepository) Create(vote *domain.Vote) error {
	r.Votes = append(r.Votes, vote)
	return nil
}

func (r *InMemoryVoteRepository) GetByID(ID string) (*domain.Vote, error) {
	for i := range r.Votes {
		if r.Votes[i].ID == ID {
			return r.Votes[i], nil
		}
	}

	return nil, nil
}

func (r *InMemoryVoteRepository) GetResults(pollID string) (map[string]uint, error) {
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
