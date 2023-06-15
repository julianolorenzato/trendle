package poll

import (
	"fmt"
	"time"

	"github.com/julianolorenzato/choosely/domain/vote"
)

type PollService struct {
	PollRepo PollRepository
	VoteRepo vote.VoteRepository
}

func NewPollService(PollRepo PollRepository, VoteRepo vote.VoteRepository) *PollService {
	return &PollService{
		PollRepo,
		VoteRepo,
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

	err = s.PollRepo.Save(p)
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

func (s *PollService) VoteInPoll(dto VoteInPollDTO) error {
	p, err := s.PollRepo.GetByID(dto.PollID)
	if err != nil {
		return err
	}

	if p == nil {
		err := fmt.Errorf("poll of id %s not found", dto.PollID)
		return err
	}

	v := vote.New(dto.VoterID, dto.PollID, dto.ChoosenOptions)

	err = p.CheckVote(v)
	if err != nil {
		return err
	}

	err = s.VoteRepo.Create(v)
	if err != nil {
		return err
	}

	return nil
}

type GetPollResultsDTO struct {
	PollID string
}

func (s *PollService) GetPollResults(dto GetPollResultsDTO) (map[string]uint, error) {
	exists := s.PollRepo.Exists(dto.PollID)
	if !exists {
		err := fmt.Errorf("poll of id %s does not exists", dto.PollID)
		return nil, err
	}

	res := s.VoteRepo.GetPollResults(dto.PollID)

	return res, nil
}
