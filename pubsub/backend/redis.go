package main

import (
	"fmt"
	"time"

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

func fetchAll(limit int) ([]FetchRoom, error) {
	var fRooms []FetchRoom

	// this call gets messages from redis
	start := time.Now()
	messages, err := fetchMessages()
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		// data is not in redis, check postgres
		fmt.Println("Redis is empty, looking into postgres...")
		start = time.Now()
		messages = RetrieveAllMessages(limit)
		fmt.Printf("Saving messages from postgres to redis...\n")
		t := time.Now()
		fmt.Printf("POSTGRES: %d items, %d milliseconds\n", len(messages), t.Sub(start).Milliseconds())
		for _, m := range messages {
			m.save()
		}
	} else {
		t := time.Now()
		fmt.Printf("REDIS: %d items, %d milliseconds\n", len(messages), t.Sub(start).Milliseconds())
	}

	for _, m := range messages {
		rInFRooms := -1
		for i, fr := range fRooms {
			if fr.Room == m.Room {
				rInFRooms = i
				break
			}
		}
		if rInFRooms == -1 {
			fRooms = append(fRooms, FetchRoom{
				Room: m.Room,
				Msgs: []Msg{},
			})
			rInFRooms = len(fRooms) - 1
		}
		fRooms[rInFRooms].Msgs = append(fRooms[rInFRooms].Msgs, Msg{
			Sender: m.Sender,
			Date:   m.Date,
			Text:   m.Text,
		})
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
	mess, err := m.MarshalBinary()
	if err != nil {
		return err
	}

	if _, err := client.RPush("messages", mess).Result(); err != nil {
		return err
	}

	return nil
}

func (r Room) save() error {
	room, err := r.MarshalBinary()
	if err != nil {
		return err
	}

	if _, err := client.RPush("rooms", room).Result(); err != nil {
		return err
	}

	return nil
}
