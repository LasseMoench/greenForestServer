package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
)

type Rooms struct {
	rooms []*Room
}

func InitRooms() *Rooms {
	return &Rooms{}
}

func (rs *Rooms) JoinAnyRoom(c *Client) {
	for _, room := range rs.rooms {
		if len(room.Clients) < 2 {
			room.addClient(c)
			c.addRoom(room)
			if len(room.Clients) == 2 {
				room.startGame()
			}
			return
		}

	}

	room := rs.NewRoom()
	room.addClient(c)
	return
}

func (rs *Rooms) NewRoom() *Room {
	id := uuid.New()
	room := &Room{
		ID:      id,
		Clients: make([]*Client, 0, 2), // Assuming max 2 clients per room
	}
	rs.rooms = append(rs.rooms, room)
	return room
}

type Room struct {
	ID      uuid.UUID
	Clients []*Client
}

const (
	Waiting int = 0
	Running     = 1
)

func (r *Room) addClient(c *Client) {
	r.Clients = append(r.Clients, c)

	for _, client := range r.Clients {
		log.Printf("Player %s joined room %s!", c.ID, r.ID)
		err := client.Conn.WriteMessage(1, []byte(fmt.Sprintf("Player %s joined room %s!", c.ID, r.ID)))
		if err != nil {
			fmt.Errorf("error sending message: %w", err)
		}
	}
}

func (r *Room) broadcast(message string) {
	for _, client := range r.Clients {
		err := client.Conn.WriteMessage(1, []byte(message))
		if err != nil {
			panic(err)
		}
	}
	return
}

func (r *Room) startGame() {
	log.Printf("Game in room %s is starting...", r.ID)
	r.broadcast("start")
}
