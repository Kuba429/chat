package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Init() {
	PORT := os.Getenv("PORT")
	if len(PORT) < 1 {
		PORT = "3000"
	}
	fmt.Printf("port: %s \n", PORT)
	log.Println("server init")
	http.HandleFunc("/ws", HandleSocket)
	http.HandleFunc("/getUserCount", func(w http.ResponseWriter, r *http.Request) {
		sum := 0
		for _, users := range MainHub.Rooms {
			sum += int(len(users))
		}
		w.Write([]byte(strconv.Itoa(sum)))
	})
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
