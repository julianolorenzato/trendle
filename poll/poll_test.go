package poll_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/julianolorenzato/choosely/poll"
	"github.com/stretchr/testify/suite"
)

type PollTestSuite struct {
	suite.Suite
	question string
	options  []string
	poll     poll.Poll
}

func (s *PollTestSuite) SetupTest() {
	s.question = "Who is the best Valorant player of all time?"
	s.options = []string{"Aspas", "Less", "Saadhak", "Cauanzin", "Tuyz"}
	p, _ := poll.NewPoll(s.question, s.options)
	s.poll = *p
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

	// --------------------------------------------------------

	p, err = poll.NewPoll("Z", s.options)

	s.Nil(p)
	s.NotNil(err)
	s.ErrorContains(err, "poll question must have at least 2 characters")

	// --------------------------------------------------------

	p, err = poll.NewPoll(strings.Repeat("Z", 51), s.options)

	s.Nil(p)
	s.NotNil(err)
	s.ErrorContains(err, "poll question must have a maximum of 50 characteres")

	// --------------------------------------------------------

	p, err = poll.NewPoll(s.question, []string{"Derke"})

	s.Nil(p)
	s.NotNil(err)
	s.ErrorContains(err, "poll must have at least 2 option")

	// --------------------------------------------------------

	opts := make([]string, 101)

	for i := range s.options {
		s.options[i] = "option" + strconv.Itoa(i)
	}

	p, err = poll.NewPoll(s.question, opts)

	s.Nil(p)
	s.NotNil(err)
	s.ErrorContains(err, "poll must have a maximum of 100 options")
}
