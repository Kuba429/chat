package server

type Message struct {
	Type       string
	Data       string
	Room       string
	SenderId   string
	SenderName string
	Date       uint64
}
