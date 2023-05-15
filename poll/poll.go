package poll

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/julianolorenzato/choosely/shared"
)

type Poll struct {
	ID              string
	Question        string
	NumberOfChoices uint32
	Options         Options
	Votes           []Vote
	IsPermanent     bool
	ExpiresAt       time.Time
	CreatedAt       time.Time
}

type Options map[string]bool

type Vote struct {
	ID             string
	VoterID        string
	OptionsChoosed []string
	CreatedAt      time.Time
}

// Enquete com ranking, onde é possível votar em mais de uma opção, e atribuir uma ordem de pesos ex: Aspas: 5, Less: 4, Saadhak: 3, Cauanzin: 2, Tuyz: 1

func NewPoll(qtn string, opts []string, nCh uint32, isPerm bool, exp time.Time) (*Poll, error) {
	if len(qtn) < 2 || len(qtn) > 50 {
		return nil, &shared.RangeError{Name: "poll question characters length", Min: 2, Max: 50}
	}

	if len(opts) < 2 || len(opts) > 100 {
		return nil, &shared.RangeError{Name: "poll options", Min: 2, Max: 100}
	}

	if nCh == 0 || nCh > uint32(len(opts)) {
		return nil, errors.New("number of choices cant be zero or larger than number of options")
	}

	if isPerm {
		exp = time.Time{}
	} else {
		if exp.Before(time.Now()) {
			return nil, &shared.ExpiredError{Name: "poll", ExpiredDate: exp}
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

// func (p *Poll) GetNumOfPossibleChoices() uint {
// 	return uint(len(p.ChoicesWeight))
// }

// func (p *Poll) Results() map[string]uint32 {
// 	result := make(map[string]uint32)

// 	for k, v := range p.Options {
// 		fmt.Println(k, "opt")
// 		result[k] = uint32(len(v))
// 	}

// 	return result
// }

// func (p *Poll) Vote(voterID string, optName string) error {
// 	_, exists := p.Options[optName]
// 	if !exists {
// 		return &shared.DoesNotExistsError{Class: "option", Name: optName}
// 	}

// 	for _, v := range p.Options {
// 		for i := range v {
// 			if v[i].VoterID == voterID {
// 				return errors.New("voter " + voterID + " already voted in this poll")
// 			}
// 		}
// 	}

// 	p.Options[optName] = append(p.Options[optName], Vote{uuid.NewString(), voterID, time.Now()})

// 	return nil
// }
