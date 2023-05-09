package poll_test

import (
	"testing"

	"github.com/julianolorenzato/choosely/poll"
	"github.com/stretchr/testify/assert"
)

func TestNewPoll(t *testing.T) {
	question := "Who is the best Valorant player of all time?"
	options := []string{"Aspas", "Less", "Saadhak", "Cauanzin", "Tuyz"}
	p, err := poll.NewPoll(question, options)

	// Poll should not be nil
	assert.NotNil(t, p)

	// NewPoll should not return a error
	assert.Nil(t, err)

	// Poll Question should be question
	assert.Equal(t, p.Question, question)

	// Poll ID should be have 36 characters (uuid)
	assert.Len(t, p.ID, 36)

	// Poll Options map should not be empty
	assert.NotEmpty(t, p.Options)

	// Poll Options map len should be equal options len
	assert.Len(t, p.Options, len(options))

	// ---------------------------------------------------

	question = "W"
	options = []string{"Aspas", "Less", "Saadhak", "Cauanzin", "Tuyz"}
	p, err = poll.NewPoll(question, options)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "poll question must have at least 2 characters")

	// ---------------------------------------------------

	question = "Who is the best Valorant player of all time?"
	options = []string{}
	p, err = poll.NewPoll(question, options)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "poll must have at least 1 option")
}
