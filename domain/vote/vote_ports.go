package vote

type VoteRepository interface {
	GetByID(ID string) *Vote
	Create(*Vote) error
}
