package connection

import (
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// This code adapted from https://outcrawl.com/realtime-collaborative-drawing-go/

type Client struct {
	id string

	hub *Hub

	// Websocket connection
	socket *websocket.Conn

	// Buffered channel of outbound messages
	outbound chan []byte
}

// Client constructor
func newClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client{
		id:       uuid.NewV4().String(),
		hub:      hub,
		socket:   socket,
		outbound: make(chan []byte),
	}
}

// Read a message from the client and send it to the hub
func (client *Client) read() {
	// When the client disconnections, unregister them
	defer func() {
		client.hub.unregister <- client
	}()

	for {
		_, data, err := client.socket.ReadMessage()
		if err != nil {
			break
		}

		client.hub.onMessage(data, client)
	}
}

// Send messages to the client from the outbound channel
func (client *Client) write() {
	for {
		select {
		case data, ok := <-client.outbound:
			if !ok {
				client.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			client.socket.WriteMessage(websocket.TextMessage, data)
		}
	}
}

func (client Client) run() {
	go client.read()
	go client.write()
}

func (client Client) close() {
	client.socket.Close()
	close(client.outbound)
}
