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
	message := Message{}
	defer func() {
		MainHub.Leave(message)
		conn.Close()
	}()
	for {
		err := conn.ReadJSON(&message)
		if err != nil {
			MainHub.Leave(message)
			break
		}
		switch message.Type {
		case "message":
			MainHub.Send(message)
		case "leave":
			MainHub.Leave(message)
		case "join":
			MainHub.Join(message, conn)
		case "username_update":
			MainHub.UpdateUsername(message)
		}
	}

}

// func printRooms(rooms map[string][]User) {
// 	fmt.Println("Rooms: ")
// 	for k := range rooms {
// 		fmt.Printf("\"%s\": %o \n", k, len(rooms[k]))
// 	}
// }
