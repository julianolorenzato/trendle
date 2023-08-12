package domain

// PollDBPort is the same as a Repository
type PollDB interface {
	GetByID(ID string) (*Poll, error)
	Create(poll *Poll) error
	Save(poll *Poll) error
}

type PollQueueProducer interface {
	notifyNewVote()
}
