package vote

import (
	"time"

	"github.com/google/uuid"
)

type Vote struct {
	ID             string
	VoterID        string
	OptionsChoosed []string
	CreatedAt      time.Time
}

func New(voterID string, optionsChoosed []string) *Vote {
	return &Vote{
		ID:             uuid.NewString(),
		VoterID:        voterID,
		OptionsChoosed: optionsChoosed,
		CreatedAt:      time.Now(),
	}
}
