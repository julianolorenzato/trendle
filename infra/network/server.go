package network

import (
	"net/http"
)

func StartServer() {
	http.HandleFunc("/ws", handleWebSocket)

	http.HandleFunc("/vote", handleVote)
	http.HandleFunc("/createPoll", handleCreatePoll)

	http.ListenAndServe("localhost:3005", nil)
}
