package core_test

import (
	"github.com/julianolorenzato/choosely/core"
	"github.com/julianolorenzato/choosely/core/domain"
	"github.com/julianolorenzato/choosely/persistence"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	sut *core.Core
}

func Test(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSubTest() {
	s.sut = &core.Core{
		PollDB: persistence.NewInMemoryPollRepository(),
		VoteDB: persistence.NewInMemoryVoteRepository(),
	}
}

func (s *TestSuite) TestCreateNewPoll() {
	s.Run("Should create a new poll", func() {
		dto := core.CreateNewPollDTO{
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
	p := &domain.Poll{
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
		s.sut.PollDB.Save(p)

		dto := core.VoteInPollDTO{
			PollID:         p.ID,
			VoterID:        uuid.NewString(),
			ChoosenOptions: []string{"Strawberry", "Orange", "Avocado"},
		}

		err := s.sut.VoteInPoll(dto)

		s.Nil(err)
	})

	s.Run("It should not vote in a poll that does not exists", func() {
		s.sut.PollDB.Save(p)

		dto := core.VoteInPollDTO{
			PollID:         uuid.NewString(), // <-- Random pollID
			VoterID:        uuid.NewString(),
			ChoosenOptions: []string{"Strawberry", "Orange", "Avocado"},
		}

		err := s.sut.VoteInPoll(dto)

		s.ErrorContains(err, "not found")
	})
}

// Only query, so does not need test?
func (s *TestSuite) TestGetPollResults() {
	s.Run("It should get the poll results", func() {
		// ...
	})
}
