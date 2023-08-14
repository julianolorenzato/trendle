package network

import (
	"encoding/json"
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

// WSClients maps [WS connection]PollID
//var wsClients = make(map[*websocket.Conn]string)

func (h *Handler) GetPollResults(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var dto core.GetPollResultsDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results, err := h.core.GetPollResults(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	// Upgrade the HTTP connection to WS connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connected websocket")
	defer conn.Close()

	// maybe i dont need this next line???
	//wsClients[conn] = dto.PollID

	// Write poll results to first message
	err = conn.WriteJSON(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Subscribe to channel and write new results every time
	// a new vote is computed in this poll
	for {
		wsW, err := conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Fatal(err)
		}

		// Pass the writer to QueueConsumer write the fresh results and close every time
		h.core.QueueConsumer.SubscribeToPollChannel(dto.PollID, wsW)
	}
}
