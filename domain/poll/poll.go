package poll

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/julianolorenzato/choosely/domain"
)

type Poll struct {
	ID              string    `json:"id"`
	Question        string    `json:"question"`
	NumberOfChoices uint32    `json:"number_of_choices"`
	Options         Options   `json:"options"`
	Votes           []Vote    `json:"votes"`
	IsPermanent     bool      `json:"is_permanent"`
	ExpiresAt       time.Time `json:"expires_at"`
	CreatedAt       time.Time `json:"created_at"`
}

type Options map[string]bool

func (o Options) exists(optName string) bool {
	_, ok := o[optName]

	return ok
}

type Vote struct {
	ID             string
	VoterID        string
	OptionsChoosed []string
	CreatedAt      time.Time
}

func NewPoll(qtn string, opts []string, nCh uint32, isPerm bool, exp time.Time) (*Poll, error) {
	if len(qtn) < 2 || len(qtn) > 50 {
		return nil, &domain.RangeError{Name: "poll question characters length", Min: 2, Max: 50}
	}

	if len(opts) < 2 || len(opts) > 100 {
		return nil, &domain.RangeError{Name: "poll options", Min: 2, Max: 100}
	}

	if nCh == 0 || nCh > uint32(len(opts)) {
		return nil, errors.New("number of choices cant be zero or larger than number of options")
	}

	if isPerm {
		exp = time.Time{}
	} else {
		if exp.Before(time.Now()) {
			return nil, &domain.ExpiredError{Name: "poll", ExpiredDate: exp}
		}
	}

	p := &Poll{
		ID:              uuid.NewString(),
		Question:        qtn,
		NumberOfChoices: nCh,
		Options:         make(map[string]bool),
		Votes:           make([]Vote, 0),
		IsPermanent:     isPerm,
		ExpiresAt:       exp,
		CreatedAt:       time.Now(),
	}

	for _, v := range opts {
		p.Options[v] = true
	}

	return p, nil
}

func (p *Poll) Results() map[string]int {
	r := make(map[string]int)

	for i := range p.Votes {
		l := len(p.Votes[i].OptionsChoosed)

		for _, o := range p.Votes[i].OptionsChoosed {
			r[o] += l
			l--
		}
	}

	return r
}

func (p *Poll) Vote(voterID string, options []string) error {
	if !p.IsPermanent && p.ExpiresAt.Before(time.Now()) {
		return &domain.ExpiredError{Name: "poll", ExpiredDate: p.ExpiresAt}
	}

	if len(options) != int(p.NumberOfChoices) {
		return errors.New("length of choosed options in a Vote must be equal Poll.NumberOfChoices")
	}

	v := Vote{
		ID:             uuid.NewString(),
		VoterID:        voterID,
		OptionsChoosed: make([]string, p.NumberOfChoices),
		CreatedAt:      time.Now(),
	}

	for i := range options {
		exists := p.Options.exists(options[i])

		if !exists {
			return &domain.DoesNotExistsError{Class: "option", Name: options[i]}
		}

		v.OptionsChoosed[i] = options[i]
	}

	p.Votes = append(p.Votes, v)

	return nil
}
