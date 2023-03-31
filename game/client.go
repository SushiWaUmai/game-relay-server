package game

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id      int `json:"clientId"`
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
	var msg Message
	var err error

	for msgData := range c.receive {
		err = json.Unmarshal(msgData, &msg)
		if err != nil {
			return
		}

		if msg.Targets == nil {
			err = c.socket.WriteMessage(websocket.TextMessage, msgData)
			if err != nil {
				return
			}
		}

		for t := range msg.Targets {
			if t == c.Id {
				err = c.socket.WriteMessage(websocket.TextMessage, msgData)
				if err != nil {
					return
				}
				break
			}
		}
	}
}
