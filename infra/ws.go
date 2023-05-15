package infra

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected websocket")

	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(message)

		err = conn.WriteMessage(mt, message)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

// func setupRoutes() *http.ServeMux {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Heyheyhey")
// 	})

// 	websocket.

// 	return mux
// }

func StartServer() {
	http.HandleFunc("/ws", handleWebSocket)

	http.ListenAndServe("localhost:3005", nil)
}
