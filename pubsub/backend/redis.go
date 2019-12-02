package main

import (
	"log"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func publish(msg Message) {
	m, err := msg.MarshalBinary()
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.Publish(msg.Room, m).Err(); err != nil {
		log.Fatalln(err)
	}
}

func subscribe(channel string, pool *Pool) {
	pubsub := client.Subscribe(channel)

	if _, err := pubsub.Receive(); err != nil {
		log.Fatalln(err)
	}

	ch := pubsub.Channel()

	go func() {
		for {
			msg := <-ch

			var m Message
			m.UnmarshalBinary([]byte(msg.Payload))

			pool.Broadcast <- m
		}
	}()
}
