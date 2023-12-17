package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID        uuid.UUID
	Conn      *websocket.Conn
	Room      *Room
	Hazelnuts int
	Problem   *Problem
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		ID:   uuid.New(),
		Conn: conn,
	}
}

func (c *Client) addRoom(room *Room) {
	c.Room = room
}

func (c *Client) handleMessages() {
	defer c.Conn.Close()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Fatalf("Error receiving message from client %s: %s", c.ID, err)
		}

		bm := &baseMessage{}
		err = json.Unmarshal(message, bm)
		if err != nil {
			log.Fatalf("Error unmarshaling base message: %s", err)
		}

		log.Printf("Received message from client %s: %v", c.ID, message)

		switch bm.MessageType {
		case "answer":
			// Parse answer
			am := &answerMessage{}
			err = json.Unmarshal(message, am)
			// Validate answer against client solution
			log.Printf("User %s sent answer %s, actual answer was: %s", c.ID, am.Answer, c.Problem.solution)
			correct := c.Problem.checkSolution(am.Answer)

			if correct {
				c.Hazelnuts += 10
			} else {
				c.Hazelnuts -= 10
			}
			c.updateProblem()

			rm := resultMessage{
				MessageType:   "result",
				ResultCorrect: correct,
				Problem:       c.Problem.problem,
				Hazelnuts:     c.Hazelnuts,
			}
			if err := c.Conn.WriteJSON(rm); err != nil {
				log.Fatalf("sending result to client %s: %s", c.ID, err)
			}
			c.Room.updateOpponent(c)
		}
		// Implement other logic here, like handling game actions
	}
}

func (c *Client) updateProblem() {
	if c.Hazelnuts > 100 {
		c.Problem = randomMxbProblem()
	} else if c.Hazelnuts > 50 {
		c.Problem = randomAdditionProblem(true)
	} else {
		c.Problem = randomAdditionProblem(false)
	}
}
