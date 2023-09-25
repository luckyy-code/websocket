package main

import (
	"log"
	"net/http"
	"websocket/ws"
)

func main() {
	setupApi()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupApi() {
	manager := ws.NewManager()
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.ServeWS)

}
