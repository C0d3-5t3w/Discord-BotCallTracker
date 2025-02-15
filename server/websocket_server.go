package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan BroadcastMessage)
var upgrader = websocket.Upgrader{}

type BroadcastMessage struct {
	Count   int    `json:"count"`
	Message string `json:"message"`
}

func StartWebSocketServer(port string) {
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	fmt.Printf("WebSocket server started on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg BroadcastMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("error:", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println("error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func BroadcastBotCallCount(count int, message string) {
	broadcast <- BroadcastMessage{Count: count, Message: message}
}
