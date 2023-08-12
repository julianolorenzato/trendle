package network

import (
	"github.com/gorilla/mux"
	"github.com/julianolorenzato/choosely/core"
	"github.com/julianolorenzato/choosely/persistence"
	"log"
	"net/http"
)

type HTTPServer struct {
	addr   string
	router *mux.Router
}

func NewHTTPServer(addr string) *HTTPServer {
	return &HTTPServer{
		addr:   addr,
		router: mux.NewRouter(),
	}
}

func (server *HTTPServer) setupRoutes() {
	pollDB := persistence.NewPostgresPollDB()
	voteDB := persistence.NewPostgresVoteDB()

	cr := core.NewCore(pollDB, voteDB)
	h := NewHandler(cr)

	server.router.HandleFunc("/poll/create", h.CreateNewPoll)
	server.router.HandleFunc("/poll/vote", h.VoteInPoll)
}

func (server *HTTPServer) Start() {
	server.setupRoutes()

	log.Println("Starting server...")

	err := http.ListenAndServe(server.addr, server.router)
	if err != nil {
		log.Fatal(err)
	}
}
