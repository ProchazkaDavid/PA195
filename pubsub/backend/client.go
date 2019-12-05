package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// Client represents websocket connection
type Client struct {
	ID     string
	Sender string
	Conn   *websocket.Conn
	Pool   *Pool
}

func (c *Client) listen() error {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	db, err := GetDBConnection()
	if err != nil {
		return err
	}

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			return err
		}

		m := map[string]string{}
		if err := json.Unmarshal(p, &m); err != nil {
			return err
		}

		switch m["event"] {
		case "send_msg":
			mess := Message{Sender: m["sender"], Text: m["text"], Date: m["date"], Room: m["room"]}
			if err := mess.save(); err != nil {
				return err
			}

			if err := InsertMessage(db, &mess); err != nil {
				return err
			}

			c.Pool.Broadcast <- Event{
				Event:  "send_msg",
				Sender: mess.Sender,
				Data:   mess,
			}
		default:
			log.Println("ERROR: JSON received in the websocket is either malformed or incorrect:", m)
		}
	}
}
