package poll

import (
	"fmt"
	"time"
)

type PollService struct {
	Repo PollRepository
}

func NewPollService(repo PollRepository) *PollService {
	return &PollService{Repo: repo}
}

type VoteInPollDTO struct {
	PollID         string
	VoterID        string
	OptionsChoosed []string
}

func (s *PollService) VoteInPoll(dto VoteInPollDTO) error {
	p, err := s.Repo.GetByID(dto.PollID)
	if err != nil {
		return err
	}
	if p == nil {
		return fmt.Errorf("poll of id %s not found", dto.PollID)
	}

	err = p.Vote(dto.VoterID, dto.OptionsChoosed)
	if err != nil {
		return err
	}

	err = s.Repo.Save(p)
	if err != nil {
		return err
	}

	return nil
}

type CreateNewPollDTO struct {
	Question        string
	Options         []string
	NumberOfChoices uint32
	IsPermanent     bool
	ExpiresAt       time.Time
}

func (s *PollService) CreateNewPoll(dto CreateNewPollDTO) error {
	p, err := NewPoll(
		dto.Question,
		dto.Options,
		dto.NumberOfChoices,
		dto.IsPermanent,
		dto.ExpiresAt,
	)
	if err != nil {
		return err
	}

	err = s.Repo.Save(p)
	if err != nil {
		return err
	}

	return nil
}
