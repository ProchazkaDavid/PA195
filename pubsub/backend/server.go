package main

import (
	"fmt"
	"net/http"
)

var channels = []string{"test"}

func serveWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := upgradeConnection(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.listen()
}

func setupRoutes(pool *Pool) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

	http.HandleFunc("/getAll", func(w http.ResponseWriter, r *http.Request) {
		// TODO: get rooms and messages from postgre to redis and then serve from redis
		// var once sync.Once
		// once.Do(func() {...})
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")

	pool := newPool()
	go pool.start()

	setupRoutes(pool)

	for _, ch := range channels {
		subscribe(ch, pool)
	}

	http.ListenAndServe(":8080", nil)
}
