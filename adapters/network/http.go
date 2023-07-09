package network

import (
	"encoding/json"
	"net/http"

	"github.com/julianolorenzato/choosely/domain/poll"
)

type PollHandler struct {
	service *poll.PollService
}

func NewPollHandler(s *poll.PollService) *PollHandler {
	return &PollHandler{service: s}
}

func (h *PollHandler) CreateNewPoll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var dto poll.CreateNewPollDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateNewPoll(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
}

func (h *PollHandler) VoteInPoll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var dto poll.VoteInPollDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.VoteInPoll(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
