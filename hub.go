package main

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
			g.Users[user] = true
		case user := <-g.Unregister:
			if _, ok := g.Users[user]; ok {
				delete(g.Users, user)
				close(user.Send)
			}
		case message := <-g.broadcast:
			for client := range g.Users {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(g.Users, client)
				}
			}
		}
	}
}