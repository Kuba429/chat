package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var hub = CreateHub()

// TODO duplicates appear when 2 connections (from 2 clients) exist. Possibly client's fault
func HandleSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	lastMessage := Message{} // store last message to leave last visited room
	for {
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Println("leave")
			hub.Leave(lastMessage)
			break
		}
		switch message.Type {
		case "message":
			fmt.Println(hub.Rooms)
		case "join":
			fmt.Println("join")
			hub.Join(message, conn)
		case "leave":
			fmt.Println("leave")
			hub.Leave(message)
			conn.Close()
		}

	}

}
