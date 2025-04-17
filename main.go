package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[string]*websocket.Conn)
var broadcast = make(chan Message)

type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Erro ao fazer upgrade:", err)
		return
	}
	defer ws.Close()

	username := r.URL.Query().Get("username")
	if username == "" {
		fmt.Println("Nome de usuário não fornecido")
		return
	}

	clients[username] = ws
	fmt.Printf("Usuário conectado: %s\n", username)

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("Erro ao ler mensagem do usuário %s: %v\n", username, err)
			delete(clients, username)
			break
		}

		msg.Sender = username
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		if recipientConn, ok := clients[msg.Recipient]; ok {
			err := recipientConn.WriteJSON(msg)
			if err != nil {
				fmt.Printf("Erro ao enviar para %s: %v\n", msg.Recipient, err)
				recipientConn.Close()
				delete(clients, msg.Recipient)
			}
		} else {
			fmt.Printf("Usuário %s não encontrado ou desconectado\n", msg.Recipient)
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	fmt.Println("Servidor iniciado na porta :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro no servidor:", err)
	}
}
