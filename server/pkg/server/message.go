package server

type Message struct {
	Type       string
	Data       string
	Image      string
	Room       string
	SenderId   string
	SenderName string
	Date       uint64
	Id         string
}
