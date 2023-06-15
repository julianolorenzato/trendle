package poll

import (
	"fmt"
	"time"

	"github.com/julianolorenzato/choosely/domain/vote"
)

type PollService struct {
	pollRepo PollRepository
	voteRepo vote.VoteRepository
}

func NewPollService(pollRepo PollRepository, voteRepo vote.VoteRepository) *PollService {
	return &PollService{
		pollRepo,
		voteRepo,
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

	err = s.pollRepo.Save(p)
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
	p, err := s.pollRepo.GetByID(dto.PollID)
	if err != nil {
		return err
	}

	if p == nil {
		err := fmt.Errorf("poll of id %s not found", dto.PollID)
		return err
	}

	v := vote.New(dto.VoterID, dto.ChoosenOptions)

	err = p.CheckVote(v)
	if err != nil {
		return err
	}

	err = s.voteRepo.Create(v)
	if err != nil {
		return err
	}

	return nil
}

type GetPollResultsDTO struct {
	PollID string
}

func (s *PollService) GetPollResults(dto GetPollResultsDTO) (map[string]uint, error) {
	exists := s.pollRepo.Exists(dto.PollID)
	if !exists {
		err := fmt.Errorf("poll of id %s does not exists", dto.PollID)
		return nil, err
	}

	res := s.voteRepo.GetPollResults(dto.PollID)

	return res, nil
}
