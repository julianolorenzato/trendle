package domain_test

import (
	"github.com/julianolorenzato/choosely/core/domain"
	"github.com/julianolorenzato/choosely/core/fail"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type PollTestSuite struct {
	suite.Suite
}

func (s *PollTestSuite) SetupSubTest() { /*...*/ }

func TestPoll(t *testing.T) {
	suite.Run(t, new(PollTestSuite))
}

func (s *PollTestSuite) TestNewPoll() {
	// Arrange
	q := "Cutest dog today?"
	o := []string{"Mary", "Pirate", "Bubbles"}

	s.Run("It should create a new Poll", func() {
		// Act
		p, err := domain.NewPoll(q, o, 3, true, time.Now())

		// Assert
		s.NotNil(p)
		s.Nil(err)
		s.Len(p.ID, 36)
		s.Equal(p.Question, q)
		s.Len(p.Options, 3)
		s.True(p.IsPermanent)
		s.EqualValues(p.NumberOfChoices, 3)
	})

	s.Run("It should not create a new Poll if the question have less than 2 characters", func() {
		// Act
		p, err := domain.NewPoll("Z", o, 3, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.IsType(err, &fail.RangeError{})
	})

	s.Run("It should not create a new Poll if the question have more than 50 characters", func() {
		// Act
		p, err := domain.NewPoll(strings.Repeat("Z", 51), o, 3, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.IsType(err, &fail.RangeError{})
	})

	s.Run("It should not create a new Poll if it have less than 2 options", func() {
		// Act
		p, err := domain.NewPoll(q, []string{"Cherry"}, 3, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.IsType(err, &fail.RangeError{})
	})

	s.Run("It should not create a new Poll if it have more than 100 options", func() {
		// Arrange
		opts := make([]string, 101)
		for i := range opts {
			opts[i] = "option" + strconv.Itoa(i)
		}

		// Act
		p, err := domain.NewPoll(q, opts, 3, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.IsType(err, &fail.RangeError{})
	})

	s.Run("It should not create a poll if the number of choices is zero", func() {
		// Act
		p, err := domain.NewPoll(q, o, 0, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.ErrorContains(err, "number of choices cant be zero or larger than number of options")
	})

	s.Run("It should not create a poll if the number of choices is greater than number of options", func() {
		// Act
		p, err := domain.NewPoll(q, []string{"Pirate", "Lucy"}, 3, true, time.Now())

		// Assert
		s.Nil(p)
		s.NotNil(err)
		s.ErrorContains(err, "number of choices cant be zero or larger than number of options")
	})

	s.Run("It should create a poll with ExpiresAt zeroed if the poll is permanent", func() {
		// Act
		p, err := domain.NewPoll(q, o, 3, true, time.Now())

		// Assert
		s.Nil(err)
		s.NotNil(p)
		s.Zero(p.ExpiresAt)
	})

	s.Run("It should create a poll with ExpiresAt not zeroed if the poll is not permanent", func() {
		// Arrange
		date := time.Now().AddDate(0, 3, 0)

		// Act
		p, err := domain.NewPoll(q, o, 3, false, date)

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
		p, err := domain.NewPoll(q, o, 3, false, date)

		// Assert
		s.NotNil(err)
		s.Nil(p)
		s.IsType(err, &fail.ExpiredError{})
	})
}

func (s *PollTestSuite) TestCheckVote() {
	s.Run("It should not return a error", func() {
		// Assert
		p := &domain.Poll{
			Options: map[string]bool{
				"first":  true,
				"second": true,
				"third":  true,
			},
			IsPermanent:     true,
			NumberOfChoices: 2,
		}

		v := &domain.Vote{
			ID:             uuid.NewString(),
			VoterID:        uuid.NewString(),
			ChoosenOptions: []string{"first", "third"},
			CreatedAt:      time.Now(),
		}

		err := p.CheckVote(v)

		s.Nil(err)
	})

	s.Run("It should not vote if the poll is not permanent and the expires date is already passed", func() {
		p := &domain.Poll{
			Options: map[string]bool{
				"first":  true,
				"second": true,
				"third":  true,
			},
			NumberOfChoices: 2,
			IsPermanent:     false,
			ExpiresAt:       time.Now().AddDate(0, 0, -1),
		}

		v := &domain.Vote{
			ID:             uuid.NewString(),
			VoterID:        uuid.NewString(),
			ChoosenOptions: []string{"first", "third"},
			CreatedAt:      time.Now(),
		}

		err := p.CheckVote(v)

		s.NotNil(err)
		s.IsType(err, &fail.ExpiredError{})
	})

	s.Run("It should not vote if the length of choosed options is different from the Poll.NumberOfChoices", func() {
		p := &domain.Poll{
			Options: map[string]bool{
				"first":  true,
				"second": true,
				"third":  true,
			},
			IsPermanent:     true,
			NumberOfChoices: 2,
		}

		v := &domain.Vote{
			ID:             uuid.NewString(),
			VoterID:        uuid.NewString(),
			ChoosenOptions: []string{"first"},
			CreatedAt:      time.Now(),
		}

		err := p.CheckVote(v)

		s.NotNil(err)
		s.ErrorContains(err, "the vote must have ")
	})

	s.Run("It should not vote if some of choosed options does not exists", func() {
		p := &domain.Poll{
			Options: map[string]bool{
				"first":  true,
				"second": true,
				"third":  true,
			},
			IsPermanent:     true,
			NumberOfChoices: 2,
		}

		v := &domain.Vote{
			ID:             uuid.NewString(),
			VoterID:        uuid.NewString(),
			ChoosenOptions: []string{"first", "fourth"},
			CreatedAt:      time.Now(),
		}

		err := p.CheckVote(v)

		s.NotNil(err)
		s.IsType(err, &fail.DoesNotExistsError{})
	})
}
