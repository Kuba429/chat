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

func HandleSocket(w http.ResponseWriter, r *http.Request) {
	var hub = CreateHub()
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println(err)
			continue
		}

		switch message.Type {
		case "message":
			fmt.Println(hub.Rooms)
		case "join":
			hub.Join(message, conn)
		case "leave":
			hub.Leave(message)
		}

	}

}
