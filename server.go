package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID     uuid.UUID
	Conn   *websocket.Conn
	Room   *Room
	Points int
}

func NewClient(id uuid.UUID, conn *websocket.Conn, room *Room) *Client {
	return &Client{
		ID:   id,
		Conn: conn,
		Room: room,
	}
}

type Room struct {
	ID      uuid.UUID
	Clients []*Client
}

type GameState struct {
	message string
}

func NewRoom(id uuid.UUID) *Room {
	return &Room{
		ID:      id,
		Clients: make([]*Client, 0, 2), // Assuming max 2 clients per room
	}
}

func joinRoom(client *Client, roomUUID uuid.UUID) {
	room, exists := rooms[roomUUID]
	if !exists {
		// Create a new room if it doesn't exist
		room = NewRoom(roomUUID)
		rooms[roomUUID] = room
	}

	// Add client to room
	room.Clients = append(room.Clients, client)

	// Check if room has enough clients to start the game
	if len(room.Clients) == 2 {
		startGame(room)
	}
}

func startGame(room *Room) {
	// Start the game with clients in the room
	// Initialize game state, notify clients, etc.

	for _, client := range room.Clients {
		err := client.Conn.WriteJSON(GameState{
			message: "start",
		})
		if err != nil {
			log.Fatal(fmt.Sprintf("error writing to client %s: %s", client.ID, err))
		}
	}
}

var (
	rooms         = make(map[uuid.UUID]*Room)
	roomJoinQueue = make(chan *Client, 10) // Adjust buffer size as needed
)

var addr = flag.String("addr", "localhost:8123", "http service address")

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	origin := r.Header["Origin"][0]
	return strings.HasPrefix(origin, "http://localhost:")
}} // use default options

func handleClientMessages(client *Client) {
	defer client.Conn.Close()

	for {
		// Read message from WebSocket
		messageType, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		// Process the message
		log.Printf("Received message from client %s: %s", client.ID, string(message))

		// For example, you can echo the message back to the client
		if err := client.Conn.WriteMessage(messageType, message); err != nil {
			log.Println("Error writing message:", err)
			break
		}

		// Implement other logic here, like handling game actions
	}
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	log.Println("Origin:", r.Header["Origin"])
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	clients := make(map[uuid.UUID]*Client)

	// Generate a unique ID for the client
	clientID := uuid.New()

	// Create a new client instance
	client := &Client{
		ID:   clientID,
		Conn: c,
	}

	// Add the client to the global list of clients
	clients[clientID] = client

	client.Conn.WriteMessage(websocket.TextMessage, []byte("Hello!"))

	for _, room := range rooms {
		if len(room.Clients) < 2 {
			joinRoom(client, room.ID)
		}
	}
	// Handle incoming messages
	go handleClientMessages(client)

}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", handleConnection)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
