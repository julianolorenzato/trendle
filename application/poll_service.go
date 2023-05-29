package application

import (
	"fmt"

	"github.com/julianolorenzato/choosely/domain/poll"
)

type PollService struct {
	Repo poll.PollRepository
}

func NewPollService(repo poll.PollRepository) *PollService {
	return &PollService{Repo: repo}
}

func (s *PollService) VoteInPoll(dto poll.VoteInPollDTO) error {
	p := s.Repo.GetByID(dto.PollID)
	if p == nil {
		return fmt.Errorf("poll of id %s not found", dto.PollID)
	}

	err := p.Vote(dto.VoterID, dto.OptionsChoosed)
	if err != nil {
		return err
	}

	err = s.Repo.Save(p)
	if err != nil {
		return err
	}

	return nil
}

func (s *PollService) CreateNewPoll(dto poll.CreateNewPollDTO) error {
	p, err := poll.NewPoll(
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
