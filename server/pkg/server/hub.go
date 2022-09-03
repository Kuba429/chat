package server

import (
	"github.com/gorilla/websocket"
)

type Hub struct {
	Rooms map[string][]User
}

func (hub Hub) Join(message Message, conn *websocket.Conn) {
	hub.Rooms[message.Room] = append(hub.Rooms[message.Room], User{
		Conn: conn,
		Id:   message.SenderId,
	})
}
func (hub Hub) Leave(message Message) {
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
}

func CreateHub() Hub {
	return Hub{Rooms: make(map[string][]User)}
}
