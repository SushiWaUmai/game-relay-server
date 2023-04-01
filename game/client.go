package game

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id      uint `json:"clientId"`
	socket  *websocket.Conn
	receive chan Message
	lobby   *Lobby
}

func (c *Client) read() {
	defer c.socket.Close()
	for {
		var msg Message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		c.lobby.forward <- msg
	}
}

func (c *Client) write() {
	defer c.socket.Close()

	for msg := range c.receive {
		msgData, err := json.Marshal(msg)
		if err != nil {
			return
		}

		err = c.socket.WriteMessage(websocket.TextMessage, msgData)
		if err != nil {
			return
		}
	}
}
