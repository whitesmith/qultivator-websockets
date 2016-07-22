package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
	"time"
)

var origin = "http://localhost/"
var url = "ws://localhost:8080/flower"

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	message := []byte(`{"sT":26.37,"eH":54.80,"eT":26.40,"lT":4036,"lI":637,"sH":16}`)
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[Sent] %s\n", message)

	time.Sleep(time.Second * 20)

	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[Sent] %s\n", message)
}