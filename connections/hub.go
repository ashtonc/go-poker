package connections

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

// This code adapted from https://outcrawl.com/realtime-collaborative-drawing-go/

type Hub struct {
	// Registered clients
	clients []*Client

	// Register requests from clients
	register chan *Client

	// Unregister requests clients
	unregister chan *Client
}

// Hub constructor
func newHub() *Hub {
	return &Hub{
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

// Get rid of this at some point
var upgrader = websocket.Upgrader{
	// Allow all origins
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (hub *Hub) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "could not upgrade", http.StatusInternalServerError)
		return
	}
	client := newClient(hub, socket)
	hub.clients = append(hub.clients, client)
	hub.register <- client
	client.run()
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


func (hub *Hub) onConnect(client *Client) {
  log.Println("client connected: ", client.socket.RemoteAddr())
  // Make list of all users
  users := []message.User{}
  for _, c := range hub.clients {
    users = append(users, message.User{ID: c.id, Name: c.Name})
  }
  // Notify user joined
  hub.send(message.NewWatcher(client.name, users), client)
  hub.broadcast(message.NewWatcherJoins(client.id, client.name), client)
}


func (hub *Hub) onJoin (client *Client){
	log.Println("client joins game:" client.socket.RemoteAddr())

}

func (hub *Hub) onDisconnect(client *Client) {
  log.Println("client disconnected: ", client.socket.RemoteAddr())
  client.close()
  // Find index of client
  i := -1
  for j, c := range hub.clients {
    if c.id == client.id {
      i = j
      break
    }
  }
  // Delete client from list
  copy(hub.clients[i:], hub.clients[i+1:])
  hub.clients[len(hub.clients)-1] = nil
  hub.clients = hub.clients[:len(hub.clients)-1]
  // Notify user left
  hub.broadcast(message.NewLeftTable(client.id), nil)
}