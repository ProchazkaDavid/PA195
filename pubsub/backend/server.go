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

func main() {
	fmt.Println("Chat App backend is running on port 8080.")
	fmt.Println("Use a command line argument to specify a limit for postgres.")

	pool := newPool()
	go pool.start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

	for _, ch := range channels {
		subscribe(ch, pool)
	}

	http.ListenAndServe(":8080", nil)
}
