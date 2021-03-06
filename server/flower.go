package main

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/whitesmith/powered-plants-web/server/models"
)

type Flower struct {
	Id string
	Conn *websocket.Conn
	Garden *Garden
	Hub *Hub
	Send chan []byte
	State []byte
}


func (flower *Flower) ReceiveMessages() {
	defer func() {
		flower.Garden.Unregister <- flower
		flower.Conn.Close()
	}()

	for {
		_, msg, err := flower.Conn.ReadMessage()
		if err != nil {
			return
		}
		payload := models.Payload{}
		json.Unmarshal([]byte(msg), &payload)

		if flower.Id != payload.Id {
			flower.Id = payload.Id
			flower.Garden.Register <- flower
		}

		flower.State = msg
		flower.Hub.broadcast <- msg
	}
}

func (flower *Flower) SendMessages()  {
	defer func() {
		flower.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-flower.Send:
			if !ok {
				flower.write(websocket.CloseMessage, []byte{})
				return
			}

			w, err := flower.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(flower.Send)
			for i := 0; i < n; i++ {
				w.Write(<-flower.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

func (flower *Flower) write(mt int, payload []byte) error {
	return flower.Conn.WriteMessage(mt, payload)
}