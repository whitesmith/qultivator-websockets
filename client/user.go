package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"
var url = "ws://localhost:8080/user"

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	var msg = make([]byte, 512)
	_, err = ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[Received] %s\n", msg)
}