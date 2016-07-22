package main

import (
	"github.com/gorilla/websocket"
	"log"
	"github.com/whitesmith/powered-plants-web/server/models"
	"encoding/json"
)

type User struct {
	Conn *websocket.Conn
	Garden *Garden
	Hub *Hub
	Send chan []byte
}

func (user *User) ReceiveMessages() {
	defer func() {
		user.Hub.Unregister <- user
		user.Conn.Close()
	}()

	for {
		_, msg, err := user.Conn.ReadMessage()
		if err != nil {
			return
		}

		command := models.Command{}
		json.Unmarshal([]byte(msg), &command)

		control := models.Control{
			Action: command.Action,
			Value: command.Value,
		}

		var request string
		if control.Value == 1{
			request = control.Action + "ON"
		}
		if control.Value == 0 {
			request = control.Action + "OFF"
		}

		log.Printf("[Flower] Sending %s", request)
		if flower, ok := user.Garden.Flowers[command.Flower]; ok {
			log.Printf("%s", request)
			flower.Send <- []byte(request)
		}
	}
}

func (user *User) SendMessages()  {
	defer func() {
		user.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-user.Send:
			log.Printf("[User] Message received: %s", message)
			if !ok {
				user.write(websocket.CloseMessage, []byte{})
				return
			}

			w, err := user.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(user.Send)
			for i := 0; i < n; i++ {
				w.Write(<-user.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

func (user *User) write(mt int, payload []byte) error {
	return user.Conn.WriteMessage(mt, payload)
}