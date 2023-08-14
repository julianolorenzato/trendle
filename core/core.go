package core

import (
	"fmt"
	"github.com/julianolorenzato/choosely/core/domain"
	"time"
)

type Core struct {
	PollDB        domain.PollDB
	VoteDB        domain.VoteDB
	QueueConsumer QueueConsumer
}

func NewCore(pollDB domain.PollDB, voteDB domain.VoteDB, qc QueueConsumer) *Core {
	return &Core{
		PollDB:        pollDB,
		VoteDB:        voteDB,
		QueueConsumer: qc,
	}
}

// ----------------------------------------------------------

type CreateNewPollDTO struct {
	Question        string
	Options         []string
	NumberOfChoices uint32
	IsPermanent     bool
	ExpiresAt       time.Time
}

func (c *Core) CreateNewPoll(dto CreateNewPollDTO) error {
	p, err := domain.NewPoll(
		dto.Question,
		dto.Options,
		dto.NumberOfChoices,
		dto.IsPermanent,
		dto.ExpiresAt,
	)
	if err != nil {
		return err
	}

	err = c.PollDB.Create(p)
	if err != nil {
		return err
	}

	return nil
}

type VoteInPollDTO struct {
	PollID         string
	VoterID        string
	ChoosenOptions []string
}

func (c *Core) VoteInPoll(dto VoteInPollDTO) error {
	p, err := c.PollDB.GetByID(dto.PollID)
	if err != nil {
		return err
	}

	if p == nil {
		err := fmt.Errorf("poll of id %s not found", dto.PollID)
		return err
	}

	v := domain.New(dto.VoterID, dto.PollID, dto.ChoosenOptions)

	err = p.CheckVote(v)
	if err != nil {
		return err
	}

	err = c.VoteDB.Create(v)
	if err != nil {
		return err
	}

	return nil
}

type GetPollResultsDTO struct {
	PollID string
}

func (c *Core) GetPollResults(dto GetPollResultsDTO) (map[string]uint, error) {
	poll, err := c.PollDB.GetByID(dto.PollID)
	if err != nil {
		return nil, err
	}

	if poll == nil {
		err := fmt.Errorf("poll of id %s does not exists", dto.PollID)
		return nil, err
	}

	res, err := c.VoteDB.GetResults(dto.PollID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
