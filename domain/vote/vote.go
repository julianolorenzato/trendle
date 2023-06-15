package vote

import (
	"time"

	"github.com/google/uuid"
)

type Vote struct {
	ID             string
	VoterID        string
	PollID         string
	ChoosenOptions []string
	CreatedAt      time.Time
}

func New(voterID string, pollID string, choosenOptions []string) *Vote {
	return &Vote{
		ID:             uuid.NewString(),
		VoterID:        voterID,
		PollID:         pollID,
		ChoosenOptions: choosenOptions,
		CreatedAt:      time.Now(),
	}
}
