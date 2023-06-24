package poll_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/julianolorenzato/choosely/adapters/persistence"
	"github.com/julianolorenzato/choosely/domain/poll"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	sut *poll.PollService
}

func Test(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSubTest() {
	s.sut = &poll.PollService{
		PollRepo: persistence.NewInMemoryPollRepository(),
		VoteRepo: persistence.NewInMemoryVoteRepository(),
	}
}

func (s *TestSuite) TestCreateNewPoll() {
	s.Run("Should create a new poll", func() {
		dto := poll.CreateNewPollDTO{
			Question:        "Some question",
			Options:         []string{"opt0", "opt1", "opt2"},
			NumberOfChoices: 2,
			IsPermanent:     true,
			ExpiresAt:       time.Now(),
		}

		err := s.sut.CreateNewPoll(dto)

		s.Nil(err)
	})
}

func (s *TestSuite) TestVoteInPoll() {
	p := &poll.Poll{
		ID:              uuid.NewString(),
		NumberOfChoices: 3,
		Options: map[string]bool{
			"Strawberry": true,
			"Orange":     true,
			"Banana":     true,
			"Apple":      true,
			"Avocado":    true,
		},
		IsPermanent: true,
	}

	s.Run("It should vote in a poll", func() {
		s.sut.PollRepo.Save(p)

		dto := poll.VoteInPollDTO{
			PollID:         p.ID,
			VoterID:        uuid.NewString(),
			ChoosenOptions: []string{"Strawberry", "Orange", "Avocado"},
		}

		err := s.sut.VoteInPoll(dto)

		s.Nil(err)
	})

	s.Run("It should not vote in a poll that does not exists", func() {
		s.sut.PollRepo.Save(p)

		dto := poll.VoteInPollDTO{
			PollID:         uuid.NewString(), // <-- Random pollID
			VoterID:        uuid.NewString(),
			ChoosenOptions: []string{"Strawberry", "Orange", "Avocado"},
		}

		err := s.sut.VoteInPoll(dto)

		s.ErrorContains(err, "not found")
	})
}

func (s *TestSuite)TestGetPollResults() {
	s.Run("It should get the poll results", func ()  {
		// ...
	})
}