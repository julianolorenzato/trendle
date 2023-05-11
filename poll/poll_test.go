package poll_test

import (
	"testing"

	"github.com/julianolorenzato/choosely/poll"
	"github.com/stretchr/testify/suite"
)

type PollTestSuite struct {
	suite.Suite
	question string
	options  []string
}

func (s *PollTestSuite) SetupTest() {
	s.question = "Who is the best Valorant player of all time?"
	s.options = []string{"Aspas", "Less", "Saadhak", "Cauanzin", "Tuyz"}
}

func TestPoll(t *testing.T) {
	suite.Run(t, new(PollTestSuite))
}

func (s *PollTestSuite) TestNewPoll() {
	p, err := poll.NewPoll(s.question, s.options)

	s.NotNil(p)
	s.Nil(err)
	s.Equal(p.Question, s.question)
	s.Len(p.ID, 36)
	s.NotEmpty(p.Options)
	s.Len(p.Options, len(s.options))

	// -------------------------------------------

	s.question = "W"
	p, err = poll.NewPoll(s.question, s.options)

	s.Nil(p)
	s.NotNil(err)
	s.ErrorContains(err, "poll question must have at least 2 characters")

	// -------------------------------------------

	s.options = []string{}
	p, err = poll.NewPoll(s.question, s.options)

	s.Nil(p)
	s.NotNil(err)
	s.ErrorContains(err, "poll question must have at least 2 characters")

	// -------------------------------------------

}
