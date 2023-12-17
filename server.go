package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var (
	addr  = flag.String("addr", "localhost:8123", "http service address")
	rooms = InitRooms()
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	origin := r.Header["Origin"][0]
	return strings.HasPrefix(origin, "http://localhost:")
}} // use default options

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", handleConnection)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	log.Println("Origin:", r.Header["Origin"])
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// Create a new client instance
	client := NewClient(c)

	// Debug
	client.Conn.WriteMessage(websocket.TextMessage, []byte("Hello!"))

	rooms.JoinAnyRoom(client)

	// Handle incoming messages
	go client.handleMessages()

}
