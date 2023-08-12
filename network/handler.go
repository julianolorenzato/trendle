package network

import (
	"encoding/json"
	"github.com/julianolorenzato/choosely/core"
	"net/http"
)

type Handler struct {
	core *core.Core
}

func NewHandler(cr *core.Core) *Handler {
	return &Handler{core: cr}
}

func (h *Handler) CreateNewPoll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var dto core.CreateNewPollDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.core.CreateNewPoll(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
}

func (h *Handler) VoteInPoll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var dto core.VoteInPollDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.core.VoteInPoll(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *Handler) GetPollResults(w http.ResponseWriter, r *http.Request) {
	handleWebSocket(w, r)
}
