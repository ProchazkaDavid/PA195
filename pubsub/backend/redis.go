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
			e := Event{
				Event:  "send_msg",
				Sender: m.Sender,
				Data:   m,
			}

			pool.Broadcast <- e
		}
	}()
}

func fetchAll() ([]FetchRoom, error) {
	var fRooms []FetchRoom

	rooms, err := fetchRooms()
	if err != nil {
		return nil, err
	}

	messages, err := fetchMessages()
	if err != nil {
		return nil, err
	}

	for _, r := range rooms {
		fRooms = append(fRooms, FetchRoom{
			Room: r.Room,
			Msgs: []Msg{},
		})
	}

	for _, m := range messages {
		for i, fr := range fRooms {
			if m.Room == fr.Room {
				fRooms[i].Msgs = append(fr.Msgs, Msg{
					Text:   m.Text,
					Sender: m.Sender,
					Date:   m.Date,
				})
			}
		}
	}

	return fRooms, nil
}

func fetchMessages() ([]Message, error) {
	var messages []Message

	mess, err := client.LRange("messages", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	for _, m := range mess {
		var message Message
		message.UnmarshalBinary([]byte(m))
		messages = append(messages, message)
	}

	return messages, nil
}

func fetchRooms() ([]Room, error) {
	var rooms []Room

	rms, err := client.LRange("rooms", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	for _, r := range rms {
		var room Room
		room.UnmarshalBinary([]byte(r))
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (m Message) save() error {
	if _, err := client.RPush("messages", m).Result(); err != nil {
		return err
	}

	return nil
}

func (r Room) save() error {
	if _, err := client.RPush("rooms", r).Result(); err != nil {
		return err
	}

	return nil
}
