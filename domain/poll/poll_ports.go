package poll

type PollRepository interface {
	GetByID(ID string) *Poll
	Create(poll *Poll) error
	Save(poll *Poll) error
}
