package main

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

func ConnectFlower(garden *Garden, hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %+v", err)
		return
	}

	flower := &Flower{Hub: hub, Garden: garden, Conn: conn, Send: make(chan []byte, 256)}
	go flower.SendMessages()
	flower.ReceiveMessages()
}

func ConnectUser(garden *Garden, hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %+v", err)
		return
	}

	user := &User{Hub: hub, Garden: garden, Conn: conn, Send: make(chan []byte, 256)}
	user.Hub.Register <- user
	go user.SendMessages()

	for _,flower := range user.Garden.Flowers {
		user.Send <- flower.State
	}
	user.ReceiveMessages()
}