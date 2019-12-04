package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func getDBConnection() *sql.DB {
	// Closing db instance left for caller. Call `defer db.Close()`
	host := os.Getenv("PG_HOST")
	port, _ := strconv.Atoi(os.Getenv("PG_PORT"))
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

type ChatMessage struct {
	Id, Sender, Time_sent, Channel, Content string
}

func retrieveAllMessages(limit int) []ChatMessage {
	db := getDBConnection()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	statement, err := db.Prepare("SELECT * FROM messages LIMIT $1")
	if err != nil {
		panic(err)
	}

	var data []ChatMessage

	rows, err := statement.Query(limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, sender, time_sent, channel, content string
		if err := rows.Scan(&id, &sender, &time_sent, &channel, &content); err != nil {
			panic(err)
		}
		row := &ChatMessage{Id: id, Sender: sender, Time_sent: time_sent, Channel: channel, Content: content}
		data = append(data, *row)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%v", data)

	return data

}
