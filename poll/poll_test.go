package poll_test

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/julianolorenzato/choosely/poll"
	"github.com/julianolorenzato/choosely/shared"
	"github.com/stretchr/testify/suite"
)

type PollTestSuite struct {
	suite.Suite
}

func (s *PollTestSuite) SetupSubTest() {
	// s.p = &poll.Poll{
	// 	ID:       "49289bb5-7228-4ee0-8a53-3ac84d3e5733",
	// 	Question: "How many brothers do you have?",
	// 	Options: map[string][]poll.Vote{
	// 		"zero": {
	// 			{"4b241f2b-295b-4ebf-b2b2-fd6ff4db20eb", "John Doe"},
	// 			{"db99a42b-a56f-42fa-970a-e81d128d6335", "Joana Doe"},
	// 		},
	// 		"one":  {},
	// 		"two":  {},
	// 		"more": {},
	// 	},
	// }
}

func TestPoll(t *testing.T) {
	suite.Run(t, new(PollTestSuite))
}

func (s *PollTestSuite) TestNewPoll() {
	// Arrange
	q := "Cutest dog today?"
	o := []string{"Mary", "Pirate", "Bubbles"}

	s.Run("It should create a new Poll", func() {
		// Act
		p, err := poll.NewPoll(q, o, 3, true, time.Now())

		// Assert
		s.NotNil(p)
		s.Nil(err)
		s.Len(p.ID, 36)
		s.Equal(p.Question, q)
		s.Len(p.Options, 3)
		s.Empty(p.Votes)
		s.True(p.IsPermanent)
		s.EqualValues(p.NumberOfChoices, 3)
	})

	s.Run("It should not create a new Poll if the question have less than 2 characters", func() {
		// Act
		p, err := poll.NewPoll("Z", o, 3, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.IsType(err, &shared.RangeError{})
	})

	s.Run("It should not create a new Poll if the question have more than 50 characters", func() {
		// Act
		p, err := poll.NewPoll(strings.Repeat("Z", 51), o, 3, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.IsType(err, &shared.RangeError{})
	})

	s.Run("It should not create a new Poll if it have less than 2 options", func() {
		// Act
		p, err := poll.NewPoll(q, []string{"Cherry"}, 3, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.IsType(err, &shared.RangeError{})
	})

	s.Run("It should not create a new Poll if it have more than 100 options", func() {
		// Arrange
		opts := make([]string, 101)
		for i := range opts {
			opts[i] = "option" + strconv.Itoa(i)
		}

		// Act
		p, err := poll.NewPoll(q, opts, 3, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.IsType(err, &shared.RangeError{})
	})

	s.Run("It should not create a poll if the number of choices is zero", func() {
		// Act
		p, err := poll.NewPoll(q, o, 0, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.ErrorContains(err, "number of choices cant be zero or larger than number of options")
	})

	s.Run("It should not create a poll if the number of choices is greater than number of options", func() {
		// Act
		p, err := poll.NewPoll(q, []string{"Pirate", "Lucy"}, 3, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.ErrorContains(err, "number of choices cant be zero or larger than number of options")
	})

	s.Run("It should create a poll with ExpiresAt zeroed if the poll is permanent", func() {
		// Act
		p, err := poll.NewPoll(q, o, 3, true, time.Now())

		// Assert
		s.Nil(err)
		s.NotNil(p)
		s.Zero(p.ExpiresAt)
	})

	s.Run("It should create a poll with ExpiresAt not zeroed if the poll is not permanent", func() {
		// Arrange
		date := time.Now().AddDate(0, 3, 0)

		// Act
		p, err := poll.NewPoll(q, o, 3, false, date)

		// Assert
		s.Nil(err)
		s.NotNil(p)
		s.NotZero(p.ExpiresAt)
		s.Equal(p.ExpiresAt, date)
	})

	s.Run("It should not create a poll if the poll is not permanent and ExpiresAt is a passed date", func() {
		// Arrange
		date := time.Now().AddDate(0, -3, 0)

		// Act
		p, err := poll.NewPoll(q, o, 3, false, date)

		// Assert
		s.NotNil(err)
		s.Nil(p)
		s.IsType(err, &shared.ExpiredError{})
	})
}

func (s *PollTestSuite) TestResults() {
	// Arrange
	p := &poll.Poll{
		ID:              "49289bb5-7228-4ee0-8a53-3ac84d3e5733",
		Question:        "Who is the GOAT actor/actress?",
		NumberOfChoices: 2,
		Options: map[string]bool{
			"Adam Sandler":   true,
			"Angelina Jolie": true,
			"Nicole Kidman":  true,
			"Will Smith":     true,
		},
		Votes: []poll.Vote{
			{OptionsChoosed: []string{"Adam Sandler", "Nicole Kidman"}},
			{OptionsChoosed: []string{"Will Smith", "Nicole Kidman"}},
			{OptionsChoosed: []string{"Angelina Jolie", "Will Smith"}},
			{OptionsChoosed: []string{"Nicole Kidman", "Angelina Jolie"}},
			{OptionsChoosed: []string{"Will Smith", "Adam Sandler"}},
		},
		IsPermanent: true,
	}

	// Act
	res := p.Results()

	// Assert
	s.EqualValues(map[string]int{
		"Adam Sandler":   3,
		"Angelina Jolie": 3,
		"Nicole Kidman":  4,
		"Will Smith":     5,
	}, res)
}

func (s *PollTestSuite) TestVote() {
	s.Run("It should vote", func() {
		p := &poll.Poll{
			Options: map[string]bool{
				"first":  true,
				"second": true,
				"third":  true,
			},
			IsPermanent: true,
			NumberOfChoices: 2,
			Votes:           []poll.Vote{},
		}

		err := p.Vote("49289bb5-7228-4ee0-8a53-3ac84d3e5733", []string{"first", "third"})

		s.Nil(err)
		s.Len(p.Votes, 1)
	})

	s.Run("It should not vote if the poll is not permanent and the expires date is already passed", func() {
		p := &poll.Poll{
			Options: map[string]bool{
				"first":  true,
				"second": true,
				"third":  true,
			},
			NumberOfChoices: 2,
			Votes:           []poll.Vote{},
			IsPermanent: false,
			ExpiresAt: time.Now().AddDate(0, 0, -1),
		}

		err := p.Vote("49289bb5-7228-4ee0-8a53-3ac84d3e5733", []string{"first", "third"})

		s.NotNil(err)
		s.IsType(err, &shared.ExpiredError{})
		s.Len(p.Votes, 0)
	})

	s.Run("It should not vote if the length of choosed options is different from the Poll.NumberOfChoices", func() {
		p := &poll.Poll{
			Options: map[string]bool{
				"first":  true,
				"second": true,
				"third":  true,
			},
			IsPermanent: true,
			NumberOfChoices: 2,
			Votes:           []poll.Vote{},
		}

		err := p.Vote("49289bb5-7228-4ee0-8a53-3ac84d3e5733", []string{"first"})

		s.NotNil(err)
		s.ErrorContains(err, "length of choosed options in a Vote must be equal Poll.NumberOfChoices")
		s.Len(p.Votes, 0)
	})

	s.Run("It should not vote if some of choosed options does not exists", func() {
		p := &poll.Poll{
			Options: map[string]bool{
				"first":  true,
				"second": true,
				"third":  true,
			},
			IsPermanent: true,
			NumberOfChoices: 2,
			Votes:           []poll.Vote{},
		}

		err := p.Vote("49289bb5-7228-4ee0-8a53-3ac84d3e5733", []string{"first", "fourth"})

		s.NotNil(err)
		s.IsType(err, &shared.DoesNotExistsError{})
		s.Len(p.Votes, 0)
	})
}
