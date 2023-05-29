package network

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julianolorenzato/choosely/domain/poll"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WSClient struct {
	Conn     *websocket.Conn
	Username string
}

type WSMessage struct {
	Action string
	User   string
	Poll   poll.Poll
}

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
