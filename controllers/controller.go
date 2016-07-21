package controllers

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/whitesmith/powered-plants-web/models"
	"github.com/whitesmith/powered-plants-web/connections"
	"github.com/whitesmith/powered-plants-web/core"
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

	connection := connections.Connection{
		Conn: conn,
	}

	for {
		message, err := connection.ReceiveMessage()
		if err != nil {
			return
		}
		log.Printf("%+v", message)
	}
	
}