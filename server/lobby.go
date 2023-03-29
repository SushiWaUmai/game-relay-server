package server

import (
	"log"
	"net/http"

	"github.com/SushiWaUmai/game-relay-server/env"
	"github.com/gorilla/websocket"
)

type Lobby struct {
	lobbyId string
	clients map[*Client]bool
	join    chan *Client
	leave   chan *Client
	forward chan []byte
}

var Lobbies = make(map[string]*Lobby)

func NewLobby() *Lobby {
	id := RandSeq(5)

	lobby := &Lobby{
		lobbyId: id,
		forward: make(chan []byte),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
	}

	Lobbies[id] = lobby
	return lobby
}

func (l *Lobby) Run() {
	for {
		select {
		case client := <-l.join:
			l.clients[client] = true
		case client := <-l.leave:
			delete(l.clients, client)
			close(client.receive)
		case msg := <-l.forward:
			for client := range l.clients {
				client.receive <- msg
			}
		}
	}
}

func (l *Lobby) PlayerNum() int {
	return len(l.clients)
}

var upgrader = &websocket.Upgrader{ReadBufferSize: env.SOCKET_BUFFER_SIZE, WriteBufferSize: env.SOCKET_BUFFER_SIZE}

func (l *Lobby) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("Failed to upgrade websoccet connection", err)
		return
	}

	client := &Client{
		socket:  socket,
		receive: make(chan []byte, env.MESSAGE_BUFFER_SIZE),
		lobby:   l,
	}

	l.join <- client
	defer func() {
		l.leave <- client
	}()
	go client.write()
	client.read()
}
