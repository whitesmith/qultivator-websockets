package main

import (
	"github.com/gorilla/websocket"
)

type User struct {
	Conn websocket.Conn
	Hub *Hub
	Garden *Garden
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
			return nil, err
		}
		//payload := Payload{}
		//json.Unmarshal([]byte(msg), &payload)
		user.Garden.broadcast <- msg
	}
}

func (user *User) SendMessages()  {
	defer func() {
		user.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-user.Send:
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