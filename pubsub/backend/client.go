package main

import (
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

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var msg Message
		msg.UnmarshalBinary(p)

		if c.Sender == "" {
			c.Sender = msg.Sender
		}

		fmt.Printf("Message Received: %+v %s \n", msg, c.ID)

		publish(msg)
		// TODO: save message to redis
		// TODO: save message to postgre
	}
}
