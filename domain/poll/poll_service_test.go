package poll_test

import (
	"testing"
	"time"

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

	// s.Run("It should not cre", func() {

	// })
}

func (s *TestSuite) TestVoteInPoll() {
	// s.Run("It should vote in a poll", func() {

	// })
}
