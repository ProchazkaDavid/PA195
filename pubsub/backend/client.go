package main

import (
	"encoding/json"
	"fmt"
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

func (c *Client) listen() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	db := GetDBConnection()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		m := map[string]string{}
		er := json.Unmarshal(p, &m)
		if er != nil {
			panic(er)
		}

		switch m["event"] {
		case "create_room":
			fmt.Printf("I should create a new room here! - room named %s\n", m["room"])
		case "send_msg":
			mess := Message{Sender: m["sender"], Text: m["text"], Date: m["date"], Room: m["room"]}
			data, _ := mess.MarshalBinary()
			fmt.Printf("marshalled: %s", data)
			err := mess.save()
			if err != nil {
				panic(err)
			}
			InsertMessage(db, &mess)
			c.Pool.Broadcast <- Event{
				Event:  "send_msg",
				Sender: mess.Sender,
				Data:   mess,
			}
		default:
			fmt.Printf("ERROR: JSON received in the websocket is either malformed or incorrect: %s\n", m)
		}

		// publish(m)
		// TODO: save message to redis
		// TODO: save message to postgre
	}
}
