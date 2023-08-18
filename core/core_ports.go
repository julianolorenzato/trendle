package core

// Abandonar output ports por entidade,
// fazer só uma port pra cada (QueueProducer, Database...)
// e colocar junto com estes input ports ??

// Input ports
type HTTPServer interface {
	Start(addr string)
}

type QueueConsumer interface {
	SubscribeToPollChannel(pollID string, callback func())
}

// Output ports
type QueueProducer interface {
	NotifyNewVote(pollID string)
}
