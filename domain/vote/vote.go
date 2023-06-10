package vote

import (
	"time"

	"github.com/google/uuid"
)

type Vote struct {
	ID             string
	VoterID        string
	ChoosenOptions []string
	CreatedAt      time.Time
}

func New(voterID string, choosenOptions []string) *Vote {
	return &Vote{
		ID:             uuid.NewString(),
		VoterID:        voterID,
		ChoosenOptions: choosenOptions,
		CreatedAt:      time.Now(),
	}
}
