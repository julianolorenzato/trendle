package poll

import (
	"errors"

	"github.com/google/uuid"
)

type Vote struct {
	ID      string
	VoterID string
}

type Poll struct {
	ID       string
	Question string
	Options  map[string][]Vote
}

func NewPoll(question string, options []string) (*Poll, error) {
	if len(question) < 2 {
		return nil, errors.New("poll question must have at least 2 characters")
	}

	if len(question) > 50 {
		return nil, errors.New("poll question must have a maximum of 50 characteres")
	}

	if len(options) < 2 {
		return nil, errors.New("poll must have at least 2 option")
	}

	if len(options) > 100 {
		return nil, errors.New("poll must have a maximum of 100 options")
	}

	p := &Poll{
		ID:       uuid.NewString(),
		Question: question,
		Options:  make(map[string][]Vote),
	}

	for _, v := range options {
		p.Options[v] = make([]Vote, 2)
	}

	return p, nil
}

func (p *Poll) AddOption(optName string) {
	p.Options[optName] = make([]Vote, 0)
}

func (p *Poll) Results() map[string]uint32 {
	result := make(map[string]uint32)

	for k, v := range p.Options {
		result[k] = uint32(len(v))
	}

	return result
}

func (p *Poll) Vote(voterID string, optName string) error {
	_, exists := p.Options[optName]
	if !exists {
		return errors.New("option " + optName + " doest not exists")
	}

	for _, v := range p.Options {
		for i := range v {
			if v[i].VoterID == voterID {
				return errors.New("voter " + voterID + " already voted in this poll")
			}
		}
	}

	p.Options[optName] = append(p.Options[optName], Vote{uuid.NewString(), voterID})

	return nil
}
