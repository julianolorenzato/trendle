package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julianolorenzato/choosely/adapters/network"
	"github.com/julianolorenzato/choosely/adapters/persistence"
	"github.com/julianolorenzato/choosely/domain/poll"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := persistence.EstablishPostgresConnection()
	if err != nil {
		log.Fatalln(err)
	}

	pollRepo := persistence.NewPostgresPollRepository(db, db)
	voteRepo := persistence.NewPostgresVoteRepository(db, db)
	pollService := poll.NewPollService(pollRepo, voteRepo)
	pollHandler := network.NewPollHandler(pollService)

	http.HandleFunc("/poll/create", pollHandler.CreateNewPoll)
	http.HandleFunc("/poll/vote", pollHandler.VoteInPoll)

	log.Println("Starting server...")

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
