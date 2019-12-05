package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// GetDBConnection establishes connection with the database.
// Closing db instance left for caller. Call `defer db.Close()`
func GetDBConnection() *sql.DB {
	host := os.Getenv("PG_HOST")
	port, _ := strconv.Atoi(os.Getenv("PG_PORT"))
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

// InsertMessage inserts a single Message struct to the database
func InsertMessage(db *sql.DB, m *Message) int {
	insertStatement := `
	INSERT INTO messages (sender, date, room, text)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err := db.QueryRow(insertStatement, m.Sender, m.Date, m.Room, m.Text).Scan(&id)
	if err != nil {
		log.Fatalln(err)
	}
	return id
}

// RetrieveAllMessages returns all the messages in a database
func RetrieveAllMessages(limit int) []Message {
	db := GetDBConnection()
	defer db.Close()

	statement, err := db.Prepare("SELECT * FROM messages LIMIT $1")
	if err != nil {
		log.Fatalln(err)
	}

	var data []Message

	rows, err := statement.Query(limit)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var m Message
		var id string
		if err := rows.Scan(&id, &m.Sender, &m.Date, &m.Room, &m.Text); err != nil {
			log.Fatalln(err)
		}
		data = append(data, m)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}

	return data

}
