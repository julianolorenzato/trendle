package network

import (
	"log"
	"net/http"
)

// NEED TO REFACTOR
func StartServer() {
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/vote", handleVote)
	http.HandleFunc("/createPoll", handleCreatePoll)

	log.Println("Starting server...")

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Error starting server: %s", err)
		}
	}()

	// Logar mensagem de sucesso
	log.Println("Server successfully started and listening on port 8080.")

	select {}
}
