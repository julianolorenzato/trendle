package poll

import "time"

type PollRepository interface {
	GetByID(ID string) *Poll
	Create(poll *Poll) error
	Save(poll *Poll) error
}

type PollService interface {
	CreateNewPoll(dto CreateNewPollDTO) error
	VoteInPoll(dto VoteInPollDTO) error
}

type CreateNewPollDTO struct {
	Question        string
	Options         []string
	NumberOfChoices uint32
	IsPermanent     bool
	ExpiresAt       time.Time
}

type VoteInPollDTO struct {
	PollID         string
	VoterID        string
	OptionsChoosed []string
}
