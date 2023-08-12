package network

import (
	"github.com/julianolorenzato/choosely/core/domain"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
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
	Poll   domain.Poll
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connected websocket")

	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		log.Println(message)

		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
