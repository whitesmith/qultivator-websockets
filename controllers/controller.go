package controllers

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ConnectDevice(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP request to Websocket
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %+v", err)
		return
	}

	// Receive messages from device
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		log.Printf("Received Message: %s", msg)
		//conn.WriteMessage(t, msg)
	}
}