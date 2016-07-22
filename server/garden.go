package main

import "log"

type Garden struct {
	Flowers map[string]*Flower
	Register chan *Flower
	Unregister chan *Flower
	broadcast chan []byte
}

func InitGarden() *Garden {
	return &Garden{
		Flowers:    make(map[string]*Flower),
		Register:   make(chan *Flower),
		Unregister: make(chan *Flower),
		broadcast:  make(chan []byte),
	}
}

func (g *Garden) run() {
	for {
		select {
		case flower := <-g.Register:
			log.Println("[Garden] Registered flower")
			g.Flowers[flower.Id] = flower
		case flower := <-g.Unregister:
			log.Println("[Garden] Unregistered flower")
			delete(g.Flowers, flower.Id)
			close(flower.Send)
		case message := <-g.broadcast:
			log.Printf("[Garden] Broadcast message: %s", message)
			for _, client := range g.Flowers {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(g.Flowers, client.Id)
				}
			}
		}
	}
}