package vote

type VoteRepository interface {
	GetResults(pollID string) (map[string]uint, error)
	// GetByID(ID string) (*Vote, error)
	Create(*Vote) error
}
