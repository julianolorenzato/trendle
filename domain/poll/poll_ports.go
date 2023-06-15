package poll

type PollRepository interface {
	Exists(ID string) bool
	GetByID(ID string) (*Poll, error)
	Create(poll *Poll) error
	Save(poll *Poll) error
}
