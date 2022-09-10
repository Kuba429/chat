package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var MainHub = CreateHub()

func HandleSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	message := Message{}
	for {
		err := conn.ReadJSON(&message)
		if err != nil {
			MainHub.Leave(message)
			break
		}
		switch message.Type {
		case "message":
			// printRooms(hub.Rooms)
			MainHub.Send(message)
		case "join":
			MainHub.Join(message, conn)
		case "leave":
			MainHub.Leave(message)
		}

	}

}

// func printRooms(rooms map[string][]User) {
// 	fmt.Println("Rooms: ")
// 	for k := range rooms {
// 		fmt.Printf("\"%s\": %o \n", k, len(rooms[k]))
// 	}
// }
