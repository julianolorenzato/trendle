package network

// import (
// 	"encoding/json"
// 	"net/http"
// 	"time"

// 	"github.com/julianolorenzato/choosely/domain/poll"
// 	"golang.org/x/exp/slices"
// )

// var polls []*poll.Poll = make([]*poll.Poll, 0)

// func handleVote(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "PUT" {
// 		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var optionsChoosed *[]string = new([]string)

// 	err := json.NewDecoder(r.Body).Decode(optionsChoosed)
// 	if err != nil {
// 		http.Error(w, err.Error()+"Error to read body", http.StatusInternalServerError)
// 		return
// 	}

// 	pollID := r.URL.Query().Get("pollID")

// 	index := slices.IndexFunc(polls, func(s *poll.Poll) bool {
// 		return s.ID == pollID
// 	})

// 	if index == -1 {
// 		http.Error(w, "Poll not found", http.StatusNotFound)
// 		return
// 	}

// 	err = polls[index].Vote("fdssfds", *optionsChoosed)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// }

// func handleCreatePoll(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
	
// 	type ReqBody struct {
// 		Question        string   `json:"question"`
// 		Options         []string `json:"options"`
// 		NumberOfChoices uint32   `json:"number_of_choices"`
// 		IsPerm          bool     `json:"is_permanent"`
// 		ExpiresInDays   int      `json:"expires_in_days"`
// 	}

// 	var body *ReqBody = new(ReqBody)

// 	err := json.NewDecoder(r.Body).Decode(body)
// 	if err != nil {
// 		http.Error(w, "Error to read the body", http.StatusBadRequest)
// 	}

// 	poll, err := poll.NewPoll(
// 		body.Question,
// 		body.Options,
// 		body.NumberOfChoices,
// 		body.IsPerm,
// 		time.Now().AddDate(0, 0, body.ExpiresInDays),
// 	)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	polls = append(polls, poll)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(poll)
// }
