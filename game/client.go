package game

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id      uint `json:"clientId"`
	socket  *websocket.Conn
	receive chan []byte
	lobby   *Lobby
}

func (c *Client) read() {
	defer c.socket.Close()
	for {
		// _, msg, err := c.socket.ReadMessage()
		var msg Message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msgData, err := json.Marshal(&msg)
		if err != nil {
			return
		}
		c.lobby.forward <- msgData
	}
}

func (c *Client) write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
