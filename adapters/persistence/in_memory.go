package persistence

import "github.com/julianolorenzato/choosely/domain/poll"

type InMemoryPollRepository struct {
	Polls []*poll.Poll
}

func NewInMemoryPollRepository() *InMemoryPollRepository {
	return &InMemoryPollRepository{
		Polls: make([]*poll.Poll, 0),
	}
}

func (r *InMemoryPollRepository) GetByID(ID string) (*poll.Poll, error){
	for i := range r.Polls {
		if r.Polls[i].ID == ID {
			return r.Polls[i], nil
		}
	}

	return nil, nil
}

func (r *InMemoryPollRepository) Create(poll *poll.Poll) error {
	r.Polls = append(r.Polls, poll)
	return nil
}

func (r *InMemoryPollRepository) Save(poll *poll.Poll) error {
	for i := range r.Polls {
		if r.Polls[i].ID == poll.ID {
			r.Polls[i] = poll
			return nil
		}
	}

	r.Polls = append(r.Polls, poll)

	return nil
}
