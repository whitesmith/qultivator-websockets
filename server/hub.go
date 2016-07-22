package main

import (
	"log"
)

type Hub struct {
	Users map[*User]bool
	Register chan *User
	Unregister chan *User
	broadcast chan []byte
}

func InitHub() *Hub {
	return &Hub{
		Users:    make(map[*User]bool),
		Register:   make(chan *User),
		Unregister: make(chan *User),
		broadcast:  make(chan []byte),
	}
}

func (g *Hub) run() {
	for {
		select {
		case user := <-g.Register:
			log.Println("[Hub] Registered user")
			g.Users[user] = true
			log.Printf("%+v", g.Users)
		case user := <-g.Unregister:
			log.Printf("%+v", g.Users)
			log.Println("[Hub] Unregistered user")
			if _, ok := g.Users[user]; ok {
				delete(g.Users, user)
				close(user.Send)
			}
		case message := <-g.broadcast:
			log.Printf("[Hub] Broadcast message: %s", message)
			for client := range g.Users {
				select {
				case client.Send <- message:
					log.Println("[Hub] Message sent")
				default:
					close(client.Send)
					delete(g.Users, client)
				}
			}
		}
	}
}