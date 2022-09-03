package server

import (
	"log"
	"net/http"
)

func Init() {
	log.Println("server init")
	http.HandleFunc("/ws", HandleSocket)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
