package connection

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"

	"poker/gamelogic"
)

func (hub *Hub) HandleWebSocket(upgrader *websocket.Upgrader, w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print(err)
		http.Error(w, "Could not upgrade.", http.StatusInternalServerError)
		return
	}

	client := newClient(hub, socket)
	hub.clients = append(hub.clients, client)
	hub.register <- client
	client.run()
	hub.send(GetGamestate(hub.game), client)
}

// This code adapted from https://outcrawl.com/realtime-collaborative-drawing-go/

type Hub struct {
	// Registered clients
	clients []*Client

	// Register requests from clients
	register chan *Client

	// Unregister requests clients
	unregister chan *Client

	game *gamelogic.Game
}

// Hub constructor
func NewHub() *Hub {
	return &Hub{
		// Create an unbuffered channel of clients
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) send(message interface{}, client *Client) {
	data, _ := json.Marshal(message)
	client.outbound <- data
}

func (hub *Hub) broadcast(message interface{}, ignore *Client) {
	data, _ := json.Marshal(message)
	for _, c := range hub.clients {
		if c != ignore {
			c.outbound <- data
		}
	}
}

func (hub *Hub) SendAll(message interface{}) {
	data, _ := json.Marshal(message)
	for _, c := range hub.clients {
		c.outbound <- data
	}
}

func (hub *Hub) onConnect(client *Client) {
	log.Print(client.socket.RemoteAddr(), " connected to the game.")
	hub.send(GetGamestate(hub.game), client)
}

func (hub *Hub) onDisconnect(client *Client) {

}

func (hub *Hub) onMessage(data []byte, client *Client) {
	kind := gjson.GetBytes(data, "kind").Int()

	// Call the game logic here
	switch kind {
	case KindTakeSeat:

	case KindLeaveSeat:

	case KindCheck:

	case KindBet:

	case KindCall:

	case KindFold:

	case KindDiscard:

	}
}
