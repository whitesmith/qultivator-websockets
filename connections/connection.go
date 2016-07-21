package connections

import (
	"encoding/json"
	"github.com/whitesmith/powered-plants-web/models"
	"github.com/gorilla/websocket"
	"github.com/whitesmith/powered-plants-web/core"
)

type Connection struct {
	Conn websocket.Conn
}


func (connection *Connection) ReceiveMessage() (models.Payload, error) {
	_, msg, err := connection.Conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	payload := models.Payload{}
	json.Unmarshal([]byte(msg), &payload)
	return payload, nil
}

func (connection *Connection) SendMessage(message string)  {
	connection.Conn.WriteMessage(1, message)
}

