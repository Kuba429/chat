package server

import (
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Rooms map[string][]User
	mu    sync.Mutex
}

func (hub *Hub) Send(message Message) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	for _, user := range hub.Rooms[message.Room] {
		user.Conn.WriteJSON(message)
	}
}

func (hub *Hub) Join(message Message, conn *websocket.Conn) {
	hub.mu.Lock()
	hub.Rooms[message.Room] = append(hub.Rooms[message.Room], User{
		Conn: conn,
		Name: message.SenderName,
		Id:   message.SenderId,
	})
	hub.mu.Unlock()

	message.Data = hub.GetUsernames(message.Room)

	hub.Send(message)
}

func (hub *Hub) Leave(message Message) {
	hub.mu.Lock()
	room := hub.Rooms[message.Room]
	if len(room) <= 1 {
		delete(hub.Rooms, message.Room)
	} else {
		for i := 0; i < len(room); i++ {
			if room[i].Id == message.SenderId {
				room[i] = room[len(room)-1]
				room[len(room)-1] = User{}
				hub.Rooms[message.Room] = room[:len(room)-1]
				break
			}
		}
	}
	hub.mu.Unlock()
	message.Type = "leave"
	message.Data = hub.GetUsernames(message.Room)
	hub.Send(message)
}
func (hub *Hub) UpdateUsername(message Message) {
	for idx, user := range MainHub.Rooms[message.Room] {
		if user.Id == message.SenderId {
			MainHub.Rooms[message.Room][idx].Name = message.SenderName
		}
	}
	message.Data = MainHub.GetUsernames(message.Room)
	MainHub.Send(message)

}

func (hub *Hub) GetUsernames(room string) string {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	var usernames []string
	for _, user := range hub.Rooms[room] {
		usernames = append(usernames, user.Name)
	}
	return strings.Join(usernames, ",")
}

func CreateHub() Hub {
	return Hub{Rooms: make(map[string][]User)}
}
