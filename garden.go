package main

type Garden struct {
	Flowers map[*Flower]bool
	Register chan *Flower
	Unregister chan *Flower
	broadcast chan []byte
}

func InitGarden() *Garden {
	return &Garden{
		Flowers:    make(map[*Flower]bool),
		Register:   make(chan *Flower),
		Unregister: make(chan *Flower),
		broadcast:  make(chan []byte),
	}
}

func (g *Garden) run() {
	for {
		select {
		case flower := <-g.Register:
			g.Flowers[flower] = true
		case flower := <-g.Unregister:
			if _, ok := g.Flowers[flower]; ok {
				delete(g.Flowers, flower)
				close(flower.Send)
			}
		case message := <-g.broadcast:
			for client := range g.Flowers {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(g.Flowers, client)
				}
			}
		}
	}
}