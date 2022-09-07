package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Init() {
	PORT := os.Getenv("PORT")
	if len(PORT) < 1 {
		PORT = "3000"
	}
	fmt.Printf("port: %s \n", PORT)
	log.Println("server init")
	http.HandleFunc("/ws", HandleSocket)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
