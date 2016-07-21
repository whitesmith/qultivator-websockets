package core

import (
	"github.com/whitesmith/powered-plants-web/models"
	"log"
)

func ReceiveMessage(payload models.Payload) () {
	log.Printf("[Device] %+v", payload)
	return
}
