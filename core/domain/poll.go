package domain

import (
	"errors"
	"fmt"
	"github.com/julianolorenzato/choosely/core/fail"
	"time"

	"github.com/google/uuid"
)

type Poll struct {
	ID              string    `json:"id"`
	Question        string    `json:"question"`
	NumberOfChoices uint32    `json:"number_of_choices"`
	Options         Options   `json:"options"`
	IsPermanent     bool      `json:"is_permanent"`
	ExpiresAt       time.Time `json:"expires_at"`
	CreatedAt       time.Time `json:"created_at"`
}

type Options map[string]bool

func (o Options) exists(optName string) bool {
	_, ok := o[optName]

	return ok
}

func (o Options) ToSlice() []string {
	keys := make([]string, 0, len(o))

	for key := range o {
		keys = append(keys, key)
	}

	return keys
}

func NewPoll(qtn string, opts []string, nCh uint32, isPerm bool, exp time.Time) (*Poll, error) {
	if len(qtn) < 2 || len(qtn) > 50 {
		return nil, &fail.RangeError{Name: "poll question characters length", Min: 2, Max: 50}
	}

	if len(opts) < 2 || len(opts) > 100 {
		return nil, &fail.RangeError{Name: "poll options", Min: 2, Max: 100}
	}

	if nCh == 0 || nCh > uint32(len(opts)) {
		return nil, errors.New("number of choices cant be zero or larger than number of options")
	}

	if isPerm {
		exp = time.Time{}
	} else {
		if exp.Before(time.Now()) {
			return nil, &fail.ExpiredError{Name: "poll", ExpiredDate: exp}
		}
	}

	p := &Poll{
		ID:              uuid.NewString(),
		Question:        qtn,
		NumberOfChoices: nCh,
		Options:         make(map[string]bool),
		IsPermanent:     isPerm,
		ExpiresAt:       exp,
		CreatedAt:       time.Now(),
	}

	for _, v := range opts {
		p.Options[v] = true
	}

	return p, nil
}

func (p *Poll) CheckVote(vote *Vote) error {
	if !p.IsPermanent && p.ExpiresAt.Before(time.Now()) {
		return &fail.ExpiredError{Name: "poll", ExpiredDate: p.ExpiresAt}
	}

	if len(vote.ChoosenOptions) != int(p.NumberOfChoices) {
		err := fmt.Errorf("the vote must have %d choosen options", p.NumberOfChoices)
		return err
	}

	for _, option := range vote.ChoosenOptions {
		exists := p.Options.exists(option)

		if !exists {
			return &fail.DoesNotExistsError{Class: "option", Name: option}
		}
	}

	return nil
}
