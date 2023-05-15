package infra

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define uma estrutura para armazenar os resultados da enquete
type PollResults struct {
	Option1 int `json:"option1"`
	Option2 int `json:"option2"`
	Option3 int `json:"option3"`
}

var pollResults = PollResults{}

// Define uma conexão WebSocket
var upgrader2 = websocket.Upgrader{}

// Configura um manipulador WebSocket para receber votos dos clientes
func voteHandler(w http.ResponseWriter, r *http.Request) {
	// Atualiza o resultado da enquete
	option := r.FormValue("option")
	switch option {
	case "option1":
		pollResults.Option1 += 1
	case "option2":
		pollResults.Option2 += 1
	case "option3":
		pollResults.Option3 += 1
	}

	// Envia uma mensagem com os novos resultados da enquete para todos os clientes conectados
	message := []byte(`{"option1": ` + string(pollResults.Option1) + `, "option2": ` + string(pollResults.Option2) + `, "option3": ` + string(pollResults.Option3) + `}`)
	for _, conn := range connections {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}

// Armazena uma lista de conexões WebSocket
var connections = []*websocket.Conn{}

// Configura um manipulador WebSocket para lidar com novas conexões
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader2.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Adiciona a conexão à lista de conexões
	connections = append(connections, conn)

	// Envia a mensagem atual dos resultados da enquete para o novo cliente conectado
	message := []byte(`{"option1": ` + string(pollResults.Option1) + `, "option2": ` + string(pollResults.Option2) + `, "option3": ` + string(pollResults.Option3) + `}`)
	conn.WriteMessage(websocket.TextMessage, message)
}

func main() {
	// Configura as rotas do servidor
	http.HandleFunc("/vote", voteHandler)
	http.HandleFunc("/ws", wsHandler)

	// Inicia o servidor
	log.Println("Servidor iniciado em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
