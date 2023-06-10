package vote

type VoteRepository interface {
	GetResults(pollID string) map[string]int
	GetByID(ID string) *Vote
	Create(*Vote) error
}
