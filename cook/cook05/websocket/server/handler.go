package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to upgrade connection: ", err)
		return
	}
	for {
		meesageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("failed to read message: ", err)
			return
		}
		log.Println("received from client: %#v", string(p))
		if err := conn.WriteMessage(meesageType, p); err != nil {
			log.Println("failed to write message: ", err)
			return
		}
	}
}
