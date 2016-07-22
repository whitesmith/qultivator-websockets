package main

import (
	"log"
	"net/http"
)

func main() {
	config = InitConfig()
	garden := InitGarden()
	hub := InitHub()
	go garden.run()
	go hub.run()

	log.Printf("Listening on port: " + config.HostPort)
	http.HandleFunc("/flower", func(w http.ResponseWriter, r *http.Request) {
		ConnectFlower(garden, hub, w, r)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		ConnectFlower(garden, hub, w, r)
	})
	err := http.ListenAndServe(":" + config.HostPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}