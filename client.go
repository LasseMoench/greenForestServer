package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID     uuid.UUID
	Conn   *websocket.Conn
	Room   *Room
	Points int
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
		// Read message from WebSocket
		messageType, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		// Process the message
		log.Printf("Received message from client %s: %s", c.ID, string(message))

		// For example, you can echo the message back to the client
		if err := c.Conn.WriteMessage(messageType, []byte(fmt.Sprintf("Received message: %s", message))); err != nil {
			log.Println("Error writing message:", err)
			break
		}

		// Implement other logic here, like handling game actions
	}
}
