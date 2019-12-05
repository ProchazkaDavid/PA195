package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Pool represents connection pool for websockets
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Event
}

// NewPool creates new pool
func newPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Event),
	}
}

// Start manages connection pool
func (pool *Pool) start() {
	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true

			fAll, err := fetchAll(limit)
			if err != nil {
				log.Println(err)
			}

			client.Conn.WriteJSON(Event{
				Event:  "fetch_all",
				Sender: "backend",
				Data:   fAll,
			})

			fmt.Println("Connected... Size of Connection Pool: ", len(pool.Clients))
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Disconnected... Size of Connection Pool: ", len(pool.Clients))
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to other clients in pool")
			for client := range pool.Clients {
				if client.Sender != message.Sender {
					if err := client.Conn.WriteJSON(message); err != nil {
						return
					}
				}
			}
		}
	}
}
