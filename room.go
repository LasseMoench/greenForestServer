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
	log.Printf("Player %s joined room %s!", c.ID, r.ID)
	r.broadcast(fmt.Sprintf("Player %s joined room %s!", c.ID, r.ID))
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
	for _, c := range r.Clients {
		// Generate new problem for client
		c.updateProblem()
		sm := gameStartMessage{
			MessageType: "start",
			Problem:     c.Problem.problem,
		}
		if err := c.Conn.WriteJSON(sm); err != nil {
			log.Fatalf("Error sending start message to client %s: %s", c.ID, err)
		}
	}
}

func (r *Room) updateOpponent(c *Client) {
	updateMessage := opponentUpdate{
		MessageType:       "opponentUpdate",
		OpponentHazelnuts: c.Hazelnuts,
	}

	log.Printf("Updating opponent. Client %s, Clients: %v", c.ID, r.Clients)
	// Find opponent and send them an update with clients points
	for _, opp := range r.Clients {
		log.Printf("client %s", c.ID)
		if opp.ID != c.ID {
			if err := opp.Conn.WriteJSON(updateMessage); err != nil {
				log.Fatal(err)
			}
		}
	}
}
