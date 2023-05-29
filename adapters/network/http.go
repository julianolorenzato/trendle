package network

import (
	"net/http"

	"github.com/julianolorenzato/choosely/domain/poll"
)

type PollHandler struct {
	service poll.PollService
	
}

func NewPollHandler(s poll.PollService) *PollHandler {
	return &PollHandler{service: s}
}

func (h *PollHandler) CreateNewPoll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	h.service.CreateNewPoll(poll.CreateNewPollDTO{})
}
