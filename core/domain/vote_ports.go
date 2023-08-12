package domain

type VoteDB interface {
	GetResults(pollID string) (map[string]uint, error)
	// GetByID(ID string) (*Vote, error)
	Create(*Vote) error
}
