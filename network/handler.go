package network

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/julianolorenzato/choosely/core"
	"log"
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *Handler) GetPollResults(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to WS connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connected websocket")
	defer conn.Close()

	params := mux.Vars(r)
	pollID, ok := params["pollID"]
	if !ok {
		http.Error(w, "need pollID in URL params", http.StatusBadRequest)
		return
	}

	var dto = core.GetPollResultsDTO{PollID: pollID}

	results, err := h.core.GetPollResults(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// Write poll results to first message
	err = conn.WriteJSON(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Subscribe to channel and write new results every time
	// a new vote is computed in this poll
	h.core.QueueConsumer.SubscribeToPollChannel(dto.PollID, func() {
		pollFreshResults, err := h.core.GetPollResults(dto)
		if err != nil {
			log.Fatal(err)
		}

		err = conn.WriteJSON(pollFreshResults)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
