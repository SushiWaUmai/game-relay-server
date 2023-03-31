package game

import (
	"log"
	"net/http"
	"sync"

	"github.com/SushiWaUmai/game-relay-server/env"
	"github.com/gorilla/websocket"
)

type Lobby struct {
	JoinCode   string `json:"joinCode"`
	clients    map[uint]*Client
	join       chan *Client
	leave      chan *Client
	forward    chan Message
	currentIdx uint
}

var Lobbies sync.Map

func NewLobby() *Lobby {
	joincode := RandSeq(5)

	lobby := &Lobby{
		JoinCode: joincode,
		forward:  make(chan Message),
		join:     make(chan *Client),
		leave:    make(chan *Client),
		clients:  make(map[uint]*Client),
	}

	Lobbies.Store(joincode, lobby)
	go lobby.Run()
	return lobby
}

func sendMsg(l *Lobby, msg Message) {
	if msg.Targets == nil {
		for _, client := range l.clients {
			client.receive <- msg
		}
	} else {
		for _, c := range msg.Targets {
			l.clients[c].receive <- msg
		}
	}
}

func (l *Lobby) Run() {
	for {
		select {
		case client := <-l.join:
			l.clients[client.Id] = client
			msg := Message{
				MsgType: "join",
				Data:    client,
				Targets: nil,
			}
			sendMsg(l, msg)
		case client := <-l.leave:
			delete(l.clients, client.Id)
			msg := Message{
				MsgType: "leave",
				Data:    client,
				Targets: nil,
			}
			sendMsg(l, msg)
			close(client.receive)
		case msg := <-l.forward:
			sendMsg(l, msg)
		}
	}
}

func (l *Lobby) ClientNum() int {
	return len(l.clients)
}

func (l *Lobby) Clients() []*Client {
	clients := make([]*Client, 0)
	for _, c := range l.clients {
		clients = append(clients, c)
	}

	return clients
}

func (l *Lobby) GetClient(id uint) *Client {
	return l.clients[id]
}

var upgrader = &websocket.Upgrader{ReadBufferSize: env.SOCKET_BUFFER_SIZE, WriteBufferSize: env.SOCKET_BUFFER_SIZE}

func (l *Lobby) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Failed to upgrade websocket connection", err)
		return
	}

	log.Println("Successfully upgraded websocket connection")

	client := &Client{
		Id:      l.currentIdx,
		socket:  socket,
		receive: make(chan Message, env.SOCKET_BUFFER_SIZE),
		lobby:   l,
	}
	l.currentIdx++

	l.join <- client
	defer func() {
		l.leave <- client
	}()
	go client.write()
	client.read()
}
