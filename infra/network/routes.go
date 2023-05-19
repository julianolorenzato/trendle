package network

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julianolorenzato/choosely/domain/poll"
	"github.com/julianolorenzato/choosely/infra/persistence"
)

func handleVote(w http.ResponseWriter, r *http.Request) {
	var optionsChoosed []string

	err := json.NewDecoder(r.Body).Decode(optionsChoosed)
	if err != nil {
		http.Error(w, "Error to read body", http.StatusInternalServerError)
		return
	}

	pollID := r.URL.Query().Get("pollID")
}

func handleCreatePoll(w http.ResponseWriter, r *http.Request) {
	type ReqBody struct {
		Question        string
		Options         []string
		NumberOfChoices uint32
		isPerm          bool
		expiresInDays   int
	}

	var body ReqBody

	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, "Error to read the body", http.StatusBadRequest)
	}

	poll, err := poll.NewPoll(
		body.Question,
		body.Options,
		body.NumberOfChoices,
		body.isPerm,
		time.Now().AddDate(0, 0, body.expiresInDays),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	persistence.

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(poll)
}
