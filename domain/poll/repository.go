package poll

type PollRepository interface {
	GetByID(ID string) Poll
	Create(poll Poll)
	Save(poll Poll)
}
