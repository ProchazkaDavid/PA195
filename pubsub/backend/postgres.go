package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// GetDBConnection establishes connection with the database.
// Closing db instance left for caller. Call `defer db.Close()`
func GetDBConnection() (*sql.DB, error) {
	host := os.Getenv("PG_HOST")
	port, _ := strconv.Atoi(os.Getenv("PG_PORT"))
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// InsertMessage inserts a single Message struct to the database
func InsertMessage(db *sql.DB, m *Message) error {
	insertStatement := `
	INSERT INTO messages (sender, date, room, text)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	var id int
	if err := db.QueryRow(insertStatement, m.Sender, m.Date, m.Room, m.Text).Scan(&id); err != nil {
		return err
	}

	return nil
}

// RetrieveAllMessages returns all the messages in a database
func RetrieveAllMessages(limit int) ([]Message, error) {
	db, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	statement, err := db.Prepare("SELECT * FROM messages LIMIT $1")
	if err != nil {
		return nil, err
	}

	rows, err := statement.Query(limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []Message
	for rows.Next() {
		var m Message
		var id string
		if err := rows.Scan(&id, &m.Sender, &m.Date, &m.Room, &m.Text); err != nil {
			return nil, err
		}
		data = append(data, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
