package main

import (
	// "github.com/julianolorenzato/choosely/application"
	"github.com/julianolorenzato/choosely/domain/poll"
	// "github.com/julianolorenzato/choosely/infra/persistence"
)

type App struct {
	Services Services
}

type Services struct {
	PollService poll.PollService
}

func Init() {
	// app := &App{
	// 	Services: Services{
	// 		PollService: application.NewPollService(persistence.),
	// 	}
	// }
}