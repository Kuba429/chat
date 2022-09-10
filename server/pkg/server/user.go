package server

import "github.com/gorilla/websocket"

type User struct {
	Conn *websocket.Conn
	Name string
	Id   string
}
